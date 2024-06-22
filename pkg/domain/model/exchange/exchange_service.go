//go:generate mockgen -source=./exchange_service.go -destination=./exchange_service_mock.gen.go -package=exchange
package exchange

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/event"
	"github.com/game-core/gc-server/pkg/domain/model/event/masterEvent"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeCost"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/masterExchangeItem"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
	"github.com/game-core/gc-server/pkg/domain/model/item"
)

type ExchangeService interface {
	Update(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeUpdateRequest) (*ExchangeUpdateResponse, error)
	Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeReceiveRequest) (*ExchangeReceiveResponse, error)
}

type exchangeService struct {
	itemService                       item.ItemService
	eventService                      event.EventService
	masterExchangeMysqlRepository     masterExchange.MasterExchangeMysqlRepository
	masterExchangeCostMysqlRepository masterExchangeCost.MasterExchangeCostMysqlRepository
	masterExchangeItemMysqlRepository masterExchangeItem.MasterExchangeItemMysqlRepository
	userExchangeMysqlRepository       userExchange.UserExchangeMysqlRepository
	userExchangeItemMysqlRepository   userExchangeItem.UserExchangeItemMysqlRepository
}

func NewExchangeService(
	itemService item.ItemService,
	eventService event.EventService,
	masterExchangeMysqlRepository masterExchange.MasterExchangeMysqlRepository,
	masterExchangeCostMysqlRepository masterExchangeCost.MasterExchangeCostMysqlRepository,
	masterExchangeItemMysqlRepository masterExchangeItem.MasterExchangeItemMysqlRepository,
	userExchangeMysqlRepository userExchange.UserExchangeMysqlRepository,
	userExchangeItemMysqlRepository userExchangeItem.UserExchangeItemMysqlRepository,
) ExchangeService {
	return &exchangeService{
		itemService:                       itemService,
		eventService:                      eventService,
		masterExchangeMysqlRepository:     masterExchangeMysqlRepository,
		masterExchangeCostMysqlRepository: masterExchangeCostMysqlRepository,
		masterExchangeItemMysqlRepository: masterExchangeItemMysqlRepository,
		userExchangeMysqlRepository:       userExchangeMysqlRepository,
		userExchangeItemMysqlRepository:   userExchangeItemMysqlRepository,
	}
}

// Update 更新
func (s *exchangeService) Update(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeUpdateRequest) (*ExchangeUpdateResponse, error) {
	masterExchangeModel, err := s.masterExchangeMysqlRepository.Find(ctx, req.MasterExchangeId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterExchangeMysqlRepository.Find", err)
	}
	masterEventModel, err := s.getEvent(ctx, now, masterExchangeModel.MasterEventId)
	if err != nil {
		return nil, errors.NewMethodError("s.getEvent", err)
	}

	userExchangeModel, userExchangeItemModels, err := s.update(ctx, tx, now, req.UserId, req.MasterExchangeId, masterEventModel)
	if err != nil {
		return nil, errors.NewMethodError("s.reset", err)
	}

	return SetExchangeUpdateResponse(userExchangeModel, userExchangeItemModels), nil
}

// Receive 受け取り
func (s *exchangeService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *ExchangeReceiveRequest) (*ExchangeReceiveResponse, error) {
	masterExchangeItemModel, masterExchangeCostModels, err := s.getMasterExchangeItemModelAndMasterExchangeCostModels(ctx, now, req.MasterExchangeItemId)
	if err != nil {
		return nil, errors.NewMethodError("s.getMasterExchangeItemModelsAndMasterExchangeCostModels", err)
	}

	userExchangeModel, userExchangeItemModel, err := s.getUserExchangeModelAndUserExchangeItemModel(ctx, req.UserId, req.MasterExchangeItemId, req.Count)
	if err != nil {
		return nil, errors.NewMethodError("s.getUserExchangeModelAndUserExchangeItemModel", err)
	}

	if err := s.consume(ctx, tx, now, req.UserId, masterExchangeCostModels); err != nil {
		return nil, errors.NewMethodError("s.consume", err)
	}

	if err := s.receive(ctx, tx, now, req.UserId, masterExchangeItemModel.MasterItemId, req.Count); err != nil {
		return nil, errors.NewMethodError("s.receive", err)
	}

	updateUserExchangeItemModel, err := s.userExchangeItemMysqlRepository.Update(
		ctx,
		tx,
		userExchangeItem.SetUserExchangeItem(
			userExchangeItemModel.UserId,
			userExchangeItemModel.MasterExchangeId,
			userExchangeItemModel.MasterExchangeItemId,
			userExchangeItemModel.Count-req.Count,
		),
	)
	if err != nil {
		return nil, errors.NewMethodError("s.userExchangeItemMysqlRepository.Update", err)
	}

	return SetExchangeReceiveResponse(userExchangeModel, updateUserExchangeItemModel), nil
}

// getMasterExchangeItemModelAndMasterExchangeCostModels マスターデータを取得する
func (s *exchangeService) getMasterExchangeItemModelAndMasterExchangeCostModels(ctx context.Context, now time.Time, masterExchangeItemId int64) (*masterExchangeItem.MasterExchangeItem, masterExchangeCost.MasterExchangeCosts, error) {
	masterExchangeItemModel, err := s.masterExchangeItemMysqlRepository.Find(ctx, masterExchangeItemId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.masterExchangeItemMysqlRepository.Find", err)
	}
	masterExchangeCostModels, err := s.masterExchangeCostMysqlRepository.FindListByMasterExchangeItemId(ctx, masterExchangeItemId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.masterExchangeCostMysqlRepository.FindByMasterExchangeItemId", err)
	}

	masterExchangeModel, err := s.masterExchangeMysqlRepository.Find(ctx, masterExchangeItemModel.MasterExchangeId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.masterExchangeMysqlRepository.Find", err)
	}
	if _, err := s.getEvent(ctx, now, masterExchangeModel.MasterEventId); err != nil {
		return nil, nil, errors.NewMethodError("s.getEvent", err)
	}

	return masterExchangeItemModel, masterExchangeCostModels, nil
}

// getMasterExchangeItemModelsAndMasterExchangeCostModels ユーザーデータを取得する
func (s *exchangeService) getUserExchangeModelAndUserExchangeItemModel(ctx context.Context, userId string, masterExchangeItemId int64, count int32) (*userExchange.UserExchange, *userExchangeItem.UserExchangeItem, error) {
	userExchangeItemModel, err := s.userExchangeItemMysqlRepository.Find(ctx, userId, masterExchangeItemId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.userExchangeItemMysqlRepository.Find", err)
	}
	if count > userExchangeItemModel.Count {
		return nil, nil, errors.NewError("over limit")
	}

	userExchangeModel, err := s.userExchangeMysqlRepository.Find(ctx, userId, userExchangeItemModel.MasterExchangeId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.userExchangeMysqlRepository.Find", err)
	}

	return userExchangeModel, userExchangeItemModel, nil
}

// getEvent イベントを取得する
func (s *exchangeService) getEvent(ctx context.Context, now time.Time, masterEventId int64) (*masterEvent.MasterEvent, error) {
	eventGerResponse, err := s.eventService.Get(ctx, event.SetEventGetRequest(masterEventId))
	if err != nil {
		return nil, errors.NewMethodError("s.eventService.Get", err)
	}

	if !eventGerResponse.MasterEvent.CheckEventPeriod(now) {
		return nil, errors.NewError("outside the event period")
	}

	return eventGerResponse.MasterEvent, nil
}

// consume アイテムを消費する
func (s *exchangeService) consume(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterExchangeCostModels masterExchangeCost.MasterExchangeCosts) error {
	items := item.NewItems()
	for _, masterExchangeCostModel := range masterExchangeCostModels {
		items = append(items, item.SetItem(masterExchangeCostModel.MasterItemId, masterExchangeCostModel.Count))
	}

	if _, err := s.itemService.Consume(ctx, tx, now, item.SetItemConsumeRequest(userId, items)); err != nil {
		return errors.NewMethodError("s.itemService.Consume", err)
	}

	return nil
}

// receive アイテムを受け取る
func (s *exchangeService) receive(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterItemId int64, count int32) error {
	items := item.NewItems()
	items = append(items, item.SetItem(masterItemId, count))

	if _, err := s.itemService.Receive(ctx, tx, now, item.SetItemReceiveRequest(userId, items)); err != nil {
		return errors.NewMethodError("s.itemService.Receive", err)
	}

	return nil
}

// update 更新する
func (s *exchangeService) update(ctx context.Context, tx *gorm.DB, now time.Time, userId string, masterExchangeId int64, masterEventModel *masterEvent.MasterEvent) (*userExchange.UserExchange, userExchangeItem.UserExchangeItems, error) {
	userExchangeModel, err := s.userExchangeMysqlRepository.FindOrNil(ctx, userId, masterExchangeId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.userExchangeMysqlRepository.Find", err)
	}

	if userExchangeModel == nil {
		newUserExchangeModel, err := s.userExchangeMysqlRepository.Create(ctx, tx, userExchange.SetUserExchange(userId, masterExchangeId, userExchangeModel.CreateResetAt(now, masterEventModel.StartAt, masterEventModel.ResetHour, masterEventModel.IntervalHour)))
		if err != nil {
			return nil, nil, errors.NewMethodError("s.userExchangeMysqlRepository.Create", err)
		}
		userExchangeItemModels, err := s.createUserExchangeItems(ctx, tx, userId, masterExchangeId)
		if err != nil {
			return nil, nil, errors.NewMethodError("s.setUserExchangeItemModels", err)
		}
		return newUserExchangeModel, userExchangeItemModels, nil
	}

	userExchangeItemModels, err := s.userExchangeItemMysqlRepository.FindListByUserIdAndMasterExchangeId(ctx, userId, masterExchangeId)
	if err != nil {
		return nil, nil, errors.NewMethodError("s.userExchangeItemMysqlRepository.FindListByUserIdAndMasterExchangeId", err)
	}

	if userExchangeModel.CheckResetAt(now, masterEventModel.IntervalHour) {
		if err := s.userExchangeItemMysqlRepository.DeleteList(ctx, tx, userExchangeItemModels); err != nil {
			return nil, nil, errors.NewMethodError("s.userExchangeItemMysqlRepository.DeleteList", err)
		}
		resetUserExchangeModel, err := s.userExchangeMysqlRepository.Update(ctx, tx, userExchange.SetUserExchange(userId, masterExchangeId, userExchangeModel.CreateResetAt(now, masterEventModel.StartAt, masterEventModel.ResetHour, masterEventModel.IntervalHour)))
		if err != nil {
			return nil, nil, errors.NewMethodError("s.userExchangeMysqlRepository.Update", err)
		}
		resetUserExchangeItemModels, err := s.createUserExchangeItems(ctx, tx, userId, masterExchangeId)
		if err != nil {
			return nil, nil, errors.NewMethodError("s.createUserExchangeItemModels", err)
		}
		return resetUserExchangeModel, resetUserExchangeItemModels, nil
	}

	return userExchangeModel, userExchangeItemModels, nil
}

// createUserExchangeItems ユーザーデータを作成する
func (s *exchangeService) createUserExchangeItems(ctx context.Context, tx *gorm.DB, userId string, masterExchangeId int64) (userExchangeItem.UserExchangeItems, error) {
	masterExchangeItemModels, err := s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId(ctx, masterExchangeId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterExchangeItemMysqlRepository.FindListByMasterExchangeId", err)
	}

	userExchangeItemModels := userExchangeItem.NewUserExchangeItems()
	for _, masterExchangeItemModel := range masterExchangeItemModels {
		userExchangeItemModels = append(
			userExchangeItemModels,
			userExchangeItem.SetUserExchangeItem(
				userId,
				masterExchangeItemModel.MasterExchangeId,
				masterExchangeItemModel.MasterExchangeItemId,
				masterExchangeItemModel.Count,
			),
		)
	}

	if _, err := s.userExchangeItemMysqlRepository.CreateList(ctx, tx, userExchangeItemModels); err != nil {
		return nil, errors.NewMethodError("s.userExchangeItemMysqlRepository.CreateList", err)
	}

	return userExchangeItemModels, nil
}
