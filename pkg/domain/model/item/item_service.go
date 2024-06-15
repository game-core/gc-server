//go:generate mockgen -source=./item_service.go -destination=./item_service_mock.gen.go -package=item
package item

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/logger"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/masterItem"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type ItemService interface {
	Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *ItemReceiveRequest) (*ItemReceiveResponse, error)
	Consume(ctx context.Context, tx *gorm.DB, now time.Time, req *ItemConsumeRequest) (*ItemConsumeResponse, error)
}

type itemService struct {
	userItemBoxMysqlRepository      userItemBox.UserItemBoxMysqlRepository
	userItemBoxCloudWatchRepository userItemBox.UserItemBoxCloudWatchRepository
	masterItemMysqlRepository       masterItem.MasterItemMysqlRepository
}

func NewItemService(
	userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository,
	userItemBoxCloudWatchRepository userItemBox.UserItemBoxCloudWatchRepository,
	masterItemMysqlRepository masterItem.MasterItemMysqlRepository,
) ItemService {
	return &itemService{
		userItemBoxMysqlRepository:      userItemBoxMysqlRepository,
		userItemBoxCloudWatchRepository: userItemBoxCloudWatchRepository,
		masterItemMysqlRepository:       masterItemMysqlRepository,
	}
}

// Receive 受け取る
func (s *itemService) Receive(ctx context.Context, tx *gorm.DB, now time.Time, req *ItemReceiveRequest) (*ItemReceiveResponse, error) {
	if err := s.checkItems(ctx, req.Items); err != nil {
		return nil, errors.NewMethodError("s.checkItems", err)
	}

	userItemBoxModels, err := s.userItemBoxMysqlRepository.FindList(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userItemMysqlRepository.FindList", err)
	}

	updateUserItemBoxModels, err := s.userItemBoxMysqlRepository.UpdateList(ctx, tx, s.setReceiveUserItemBox(req.UserId, req.Items, userItemBoxModels))
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.UpdateList", err)
	}

	s.userItemBoxCloudWatchRepository.CreateList(ctx, now, logger.LogLevel_Success, updateUserItemBoxModels)
	return SetItemReceiveResponse(updateUserItemBoxModels), nil
}

// Consume 消費する
func (s *itemService) Consume(ctx context.Context, tx *gorm.DB, now time.Time, req *ItemConsumeRequest) (*ItemConsumeResponse, error) {
	if err := s.checkItems(ctx, req.Items); err != nil {
		return nil, errors.NewMethodError("s.checkItems", err)
	}

	userItemBoxModels, err := s.userItemBoxMysqlRepository.FindList(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userItemMysqlRepository.FindList", err)
	}

	updateUserItemBoxModels, err := s.setConsumeUserItemBox(req.Items, userItemBoxModels)
	if err != nil {
		return nil, errors.NewMethodError("s.setConsumeUserItemBox", err)
	}

	if _, err := s.userItemBoxMysqlRepository.UpdateList(ctx, tx, updateUserItemBoxModels); err != nil {
		return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.UpdateList", err)
	}

	s.userItemBoxCloudWatchRepository.CreateList(ctx, now, logger.LogLevel_Success, updateUserItemBoxModels)
	return SetItemConsumeResponse(updateUserItemBoxModels), nil
}

// setReceiveUserItemBox 受け取るItemBox一覧をセットする
func (s *itemService) setReceiveUserItemBox(userId string, items Items, userItemBoxModels userItemBox.UserItemBoxes) userItemBox.UserItemBoxes {
	userItemBoxMaps := userItemBoxModels.SetUserItemBoxMaps()
	updateUserItemBoxModels := userItemBox.NewUserItemBoxes()

	for _, itemModel := range items {
		userItemBoxModel, exist := userItemBoxMaps[itemModel.MasterItemId]
		if exist {
			updateUserItemBoxModels = append(updateUserItemBoxModels, userItemBox.SetUserItemBox(userItemBoxModel.UserId, userItemBoxModel.MasterItemId, userItemBoxModel.Count+itemModel.Count))
		} else {
			updateUserItemBoxModels = append(updateUserItemBoxModels, userItemBox.SetUserItemBox(userId, itemModel.MasterItemId, itemModel.Count))
		}
	}

	return updateUserItemBoxModels
}

// setConsumeUserItemBox 消費するItemBox一覧をセットする
func (s *itemService) setConsumeUserItemBox(items Items, userItemBoxModels userItemBox.UserItemBoxes) (userItemBox.UserItemBoxes, error) {
	userItemBoxMaps := userItemBoxModels.SetUserItemBoxMaps()
	updateUserItemBoxModels := userItemBox.NewUserItemBoxes()

	for _, itemModel := range items {
		userItemBoxModel, exist := userItemBoxMaps[itemModel.MasterItemId]
		if !exist {
			return nil, errors.NewError("don't have item")
		}
		count := userItemBoxModel.Count - itemModel.Count
		if count < 0 {
			return nil, errors.NewError("missing item")
		}
		updateUserItemBoxModels = append(updateUserItemBoxModels, userItemBox.SetUserItemBox(userItemBoxModel.UserId, userItemBoxModel.MasterItemId, count))
	}

	return updateUserItemBoxModels, nil
}

// checkItems アイテムが存在するか確認する
func (s *itemService) checkItems(ctx context.Context, items Items) error {
	for _, item := range items {
		if _, err := s.masterItemMysqlRepository.Find(ctx, item.MasterItemId); err != nil {
			return errors.NewMethodError("s.masterItemMysqlRepository.Find", err)
		}
	}

	return nil
}
