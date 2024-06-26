//go:generate mockgen -source=./action_service.go -destination=./action_service_mock.gen.go -package=action
package action

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterAction"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
	"github.com/game-core/gc-server/pkg/domain/model/action/userAction"
)

type ActionService interface {
	GetListMaster(ctx context.Context) (*ActionGetListMasterResponse, error)
	Check(ctx context.Context, now time.Time, req *ActionCheckRequest) error
	Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) error
}

type actionService struct {
	masterActionMysqlRepository        masterAction.MasterActionMysqlRepository
	masterActionRunMysqlRepository     masterActionRun.MasterActionRunMysqlRepository
	masterActionStepMysqlRepository    masterActionStep.MasterActionStepMysqlRepository
	masterActionTriggerMysqlRepository masterActionTrigger.MasterActionTriggerMysqlRepository
	userActionMysqlRepository          userAction.UserActionMysqlRepository
}

func NewActionService(
	masterActionMysqlRepository masterAction.MasterActionMysqlRepository,
	masterActionRunMysqlRepository masterActionRun.MasterActionRunMysqlRepository,
	masterActionStepMysqlRepository masterActionStep.MasterActionStepMysqlRepository,
	masterActionTriggerMysqlRepository masterActionTrigger.MasterActionTriggerMysqlRepository,
	userActionMysqlRepository userAction.UserActionMysqlRepository,
) ActionService {
	return &actionService{
		masterActionMysqlRepository:        masterActionMysqlRepository,
		masterActionRunMysqlRepository:     masterActionRunMysqlRepository,
		masterActionStepMysqlRepository:    masterActionStepMysqlRepository,
		masterActionTriggerMysqlRepository: masterActionTriggerMysqlRepository,
		userActionMysqlRepository:          userActionMysqlRepository,
	}
}

// GetListMaster アクション一覧を取得する
func (s *actionService) GetListMaster(ctx context.Context) (*ActionGetListMasterResponse, error) {
	masterActionModels, err := s.masterActionMysqlRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionMysqlRepository.FindList", err)
	}

	masterActionRunModels, err := s.masterActionRunMysqlRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionRunMysqlRepository.FindList", err)
	}

	masterActionStepModels, err := s.masterActionStepMysqlRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionStepMysqlRepository.FindList", err)
	}

	masterActionTriggerModels, err := s.masterActionTriggerMysqlRepository.FindList(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionTriggerMysqlRepository.FindList", err)
	}

	return SetActionGetListMasterResponse(
		masterActionModels,
		masterActionRunModels,
		masterActionStepModels,
		masterActionTriggerModels,
	), nil
}

// Check アクションが実行済か確認する
func (s *actionService) Check(ctx context.Context, now time.Time, req *ActionCheckRequest) error {
	masterActionModel, err := s.getMasterAction(ctx, req.MasterActionStepEnum, req.TargetId)
	if err != nil {
		return errors.NewMethodError("s.getMasterAction", err)
	}

	if masterActionModel == nil {
		return nil
	}

	if _, err := s.getUserAction(ctx, now, req.UserId, masterActionModel); err != nil {
		return errors.NewMethodError("s.getUserAction", err)
	}

	return nil
}

// Run アクションを実行する
func (s *actionService) Run(ctx context.Context, tx *gorm.DB, now time.Time, req *ActionRunRequest) error {
	masterActionModel, err := s.getMasterAction(ctx, req.MasterActionStepEnum, req.TargetId)
	if err != nil {
		return errors.NewMethodError("s.getMasterAction", err)
	}

	if masterActionModel == nil {
		return nil
	}

	triggerUserActionModel, triggerMasterActionModel, err := s.getTriggerUserAction(ctx, now, req.UserId, masterActionModel.TriggerMasterActionId)
	if err != nil {
		return errors.NewMethodError("s.checkTriggerUserAction", err)
	}

	if err := s.deleteTriggerUserAction(ctx, tx, triggerUserActionModel, triggerMasterActionModel); err != nil {
		return errors.NewMethodError("s.deleteTriggerUserAction", err)
	}

	if err := s.run(ctx, tx, now, req.UserId, masterActionModel); err != nil {
		return errors.NewMethodError("s.run", err)
	}

	return nil
}

// getMasterAction マスターアクションを取得する
func (s *actionService) getMasterAction(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*masterAction.MasterAction, error) {
	if targetId != nil {
		masterActionModel, err := s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId(ctx, masterActionStepEnum, targetId)
		if err != nil {
			return nil, errors.NewMethodError("s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnumAndTargetId", err)
		}

		return masterActionModel, nil
	}

	masterActionModel, err := s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum(ctx, masterActionStepEnum)
	if err != nil {
		return nil, errors.NewMethodError("s.masterActionMysqlRepository.FindOrNilByMasterActionStepEnum", err)
	}

	return masterActionModel, nil
}

// getUserAction アクションを取得する
func (s *actionService) getUserAction(ctx context.Context, now time.Time, userId string, masterActionModel *masterAction.MasterAction) (*userAction.UserAction, error) {
	userActionModel, err := s.userActionMysqlRepository.Find(ctx, userId, masterActionModel.MasterActionId)
	if err != nil {
		return nil, errors.NewMethodError("s.userActionMysqlRepository.Find", err)
	}

	if !userActionModel.CheckExpiration(now, masterActionModel.Expiration) {
		return nil, errors.NewError("expiration date has expired")
	}

	return userActionModel, nil
}

// getTriggerUserAction トリガーになるアクションを取得する
func (s *actionService) getTriggerUserAction(ctx context.Context, now time.Time, userId string, triggerActionId *int64) (*userAction.UserAction, *masterAction.MasterAction, error) {
	if triggerActionId == nil {
		return nil, nil, nil
	}

	triggerMasterActionModel, err := s.masterActionMysqlRepository.Find(ctx, *triggerActionId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.masterActionMysqlRepository.Find", err)
	}

	triggerUserActionModel, err := s.getUserAction(ctx, now, userId, triggerMasterActionModel)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.getUserAction", err)
	}

	return triggerUserActionModel, triggerMasterActionModel, nil
}

// deleteTriggerAction トリガーアクションを削除する
func (s *actionService) deleteTriggerUserAction(ctx context.Context, tx *gorm.DB, triggerUserActionModel *userAction.UserAction, triggerMasterActionModel *masterAction.MasterAction) error {
	if triggerUserActionModel != nil {
		switch triggerMasterActionModel.MasterActionTriggerEnum {
		case masterActionTrigger.MasterActionTriggerEnum_Continuation:
			return nil
		case masterActionTrigger.MasterActionTriggerEnum_Discontinuation:
			if err := s.userActionMysqlRepository.Delete(ctx, tx, triggerUserActionModel); err != nil {
				return errors.NewMethodError("s.userActionMysqlRepository.Delete", err)
			}
		default:
			return errors.NewError("MasterActionTriggerType does not exist")
		}
	}

	return nil
}

// run 実行する
func (s *actionService) run(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterActionModel *masterAction.MasterAction) error {
	if err := s.update(ctx, tx, userAction.SetUserAction(userId, masterActionModel.MasterActionId, now)); err != nil {
		return errors.NewMethodError("s.update", err)
	}

	// 実行されるアクションがある場合
	masterActionRunModels, err := s.masterActionRunMysqlRepository.FindListByMasterActionId(ctx, masterActionModel.MasterActionId)
	if err != nil {
		return errors.NewMethodError("s.masterActionRunMysqlRepository.FindListByMasterActionId", err)
	}

	for _, model := range masterActionRunModels {
		if err := s.update(ctx, tx, userAction.SetUserAction(userId, model.MasterActionId, now)); err != nil {
			return errors.NewMethodError("s.userActionMysqlRepository.Create", err)
		}
	}

	return nil
}

// update 更新する
func (s *actionService) update(ctx context.Context, tx *gorm.DB, model *userAction.UserAction) error {
	userActionModel, err := s.userActionMysqlRepository.FindOrNil(ctx, model.UserId, model.MasterActionId)
	if err != nil {
		return errors.NewMethodError("s.userActionMysqlRepository.FindOrNil", err)
	}

	if userActionModel != nil {
		if _, err := s.userActionMysqlRepository.Update(ctx, tx, model); err != nil {
			return errors.NewMethodError("s.userActionMysqlRepository.Update", err)
		}

		return nil
	}

	if _, err := s.userActionMysqlRepository.Create(ctx, tx, model); err != nil {
		return errors.NewMethodError("s.userActionMysqlRepository.Create", err)
	}

	return nil
}
