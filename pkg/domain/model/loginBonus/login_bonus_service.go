//go:generate mockgen -source=./login_bonus_service.go -destination=./login_bonus_service_mock.gen.go -package=loginBonus
package loginBonus

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/event"
	"github.com/game-core/gc-server/pkg/domain/model/event/masterEvent"
	"github.com/game-core/gc-server/pkg/domain/model/item"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonus"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusItem"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/masterLoginBonusSchedule"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusService interface {
	Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error)
}

type loginBonusService struct {
	itemService                             item.ItemService
	eventService                            event.EventService
	userLoginBonusMysqlRepository           userLoginBonus.UserLoginBonusMysqlRepository
	masterLoginBonusMysqlRepository         masterLoginBonus.MasterLoginBonusMysqlRepository
	masterLoginBonusItemMysqlRepository     masterLoginBonusItem.MasterLoginBonusItemMysqlRepository
	masterLoginBonusScheduleMysqlRepository masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository
}

func NewLoginBonusService(
	itemService item.ItemService,
	eventService event.EventService,
	userLoginBonusMysqlRepository userLoginBonus.UserLoginBonusMysqlRepository,
	masterLoginBonusMysqlRepository masterLoginBonus.MasterLoginBonusMysqlRepository,
	masterLoginBonusItemMysqlRepository masterLoginBonusItem.MasterLoginBonusItemMysqlRepository,
	masterLoginBonusScheduleMysqlRepository masterLoginBonusSchedule.MasterLoginBonusScheduleMysqlRepository,
) LoginBonusService {
	return &loginBonusService{
		itemService:                             itemService,
		eventService:                            eventService,
		userLoginBonusMysqlRepository:           userLoginBonusMysqlRepository,
		masterLoginBonusMysqlRepository:         masterLoginBonusMysqlRepository,
		masterLoginBonusItemMysqlRepository:     masterLoginBonusItemMysqlRepository,
		masterLoginBonusScheduleMysqlRepository: masterLoginBonusScheduleMysqlRepository,
	}
}

// Receive ログインボーナスを受け取る
func (s *loginBonusService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *LoginBonusReceiveRequest) (*LoginBonusReceiveResponse, error) {
	masterLoginBonusModel, err := s.masterLoginBonusMysqlRepository.Find(ctx, req.MasterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusMysqlRepository.Find", err)
	}

	masterLoginBonusEventModel, err := s.getEvent(ctx, now, masterLoginBonusModel.MasterEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	masterLoginBonusScheduleModel, err := s.getSchedule(ctx, now, req.MasterLoginBonusId, masterLoginBonusEventModel.IntervalHour, masterLoginBonusEventModel.StartAt)
	if err != nil {
		return nil, errors.NewMethodError("s.getSchedule", err)
	}

	masterLoginBonusItemModels, err := s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId(ctx, masterLoginBonusScheduleModel.MasterLoginBonusScheduleId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusItemMysqlRepository.FindListByMasterLoginBonusScheduleId", err)
	}

	userLoginBonusModel, err := s.getUser(ctx, now, req.UserId, req.MasterLoginBonusId, masterLoginBonusEventModel.ResetHour)
	if err != nil {
		return nil, errors.NewMethodError("s.getUser", err)
	}

	if err := s.receive(ctx, tx, now, req.UserId, masterLoginBonusItemModels); err != nil {
		return nil, errors.NewMethodError("s.receive", err)
	}

	result, err := s.update(ctx, tx, now, req.UserId, req.MasterLoginBonusId, userLoginBonusModel)
	if err != nil {
		return nil, errors.NewMethodError("s.update", err)
	}

	return SetLoginBonusReceiveResponse(result), nil
}

// getEvent イベントを取得する
func (s *loginBonusService) getEvent(ctx context.Context, now time.Time, masterEventId int64) (*masterEvent.MasterEvent, error) {
	eventGerResponse, err := s.eventService.Get(ctx, event.SetEventGetRequest(masterEventId))
	if err != nil {
		return nil, errors.NewMethodError("s.eventService.Get", err)
	}

	if !eventGerResponse.MasterEvent.CheckEventPeriod(now) {
		return nil, errors.NewError("outside the event period")
	}

	return eventGerResponse.MasterEvent, nil
}

// getSchedule スケジュールを取得する
func (s *loginBonusService) getSchedule(ctx context.Context, now time.Time, masterLoginBonusId int64, intervalHour int32, startAt time.Time) (*masterLoginBonusSchedule.MasterLoginBonusSchedule, error) {
	masterLoginBonusSchedules, err := s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId(ctx, masterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterLoginBonusScheduleMysqlRepository.FindListByMasterLoginBonusId", err)
	}

	return masterLoginBonusSchedules.GetScheduleByStep(masterLoginBonusSchedules.GetStep(intervalHour, startAt, now)), nil
}

// getUser 取得する
func (s *loginBonusService) getUser(ctx context.Context, now time.Time, userId string, masterLoginBonusId int64, resetHour int32) (*userLoginBonus.UserLoginBonus, error) {
	userLoginBonusModel, err := s.userLoginBonusMysqlRepository.FindOrNil(ctx, userId, masterLoginBonusId)
	if err != nil {
		return nil, errors.NewMethodError("s.userLoginBonusMysqlRepository.FindOrNil", err)
	}

	if userLoginBonusModel.CheckReceived(resetHour, now) {
		return nil, errors.NewError("already received")
	}

	return userLoginBonusModel, nil
}

// receive 受け取り
func (s *loginBonusService) receive(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterLoginBonusItemModels masterLoginBonusItem.MasterLoginBonusItems) error {
	items := item.NewItems()
	for _, masterLoginBonusItemModel := range masterLoginBonusItemModels {
		items = append(items, item.SetItem(masterLoginBonusItemModel.MasterItemId, masterLoginBonusItemModel.Count))
	}

	if _, err := s.itemService.Receive(ctx, tx, now, item.SetItemReceiveRequest(userId, items)); err != nil {
		return errors.NewMethodError("s.itemService.Receive", err)
	}

	return nil
}

// update ユーザーログインボーナスを更新
func (s *loginBonusService) update(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterLoginBonusId int64, userLoginBonusModel *userLoginBonus.UserLoginBonus) (*userLoginBonus.UserLoginBonus, error) {
	if userLoginBonusModel != nil {
		userLoginBonusModel.ReceivedAt = now
		result, err := s.userLoginBonusMysqlRepository.Update(ctx, tx, userLoginBonusModel)
		if err != nil {
			return nil, errors.NewMethodError("s.userLoginBonusMysqlRepository.Update", err)
		}

		return result, nil
	}

	result, err := s.userLoginBonusMysqlRepository.Create(ctx, tx, userLoginBonus.SetUserLoginBonus(userId, masterLoginBonusId, now))
	if err != nil {
		return nil, errors.NewMethodError("s.userLoginBonusMysqlRepository.Create", err)
	}

	return result, nil
}
