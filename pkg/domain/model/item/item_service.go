//go:generate mockgen -source=./item_service.go -destination=./item_service_mock.gen.go -package=item
package item

import (
	"context"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/masterItem"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type ItemService interface {
	Receive(ctx context.Context, tx *gorm.DB, req *ItemReceiveRequest) (*ItemReceiveResponse, error)
}

type itemService struct {
	userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository
	masterItemMysqlRepository  masterItem.MasterItemMysqlRepository
}

func NewItemService(
	userItemBoxMysqlRepository userItemBox.UserItemBoxMysqlRepository,
	masterItemMysqlRepository masterItem.MasterItemMysqlRepository,
) ItemService {
	return &itemService{
		userItemBoxMysqlRepository: userItemBoxMysqlRepository,
		masterItemMysqlRepository:  masterItemMysqlRepository,
	}
}

// Receive 受け取る
func (s *itemService) Receive(ctx context.Context, tx *gorm.DB, req *ItemReceiveRequest) (*ItemReceiveResponse, error) {
	if err := s.checkItems(ctx, req.Items); err != nil {
		return nil, errors.NewMethodError("s.checkItems", err)
	}

	userItemBoxModels, err := s.userItemBoxMysqlRepository.FindList(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userItemMysqlRepository.FindList", err)
	}

	updateUserItemBoxModels, err := s.userItemBoxMysqlRepository.UpdateList(ctx, tx, s.setUpdateUserItemBox(req.UserId, req.Items, userItemBoxModels))
	if err != nil {
		return nil, errors.NewMethodError("s.userItemBoxMysqlRepository.UpdateList", err)
	}

	return SetItemReceiveResponse(updateUserItemBoxModels), nil
}

// SetUpdateUserItemBox 更新するItemBox一覧をセットする
func (s *itemService) setUpdateUserItemBox(userId string, items Items, userItemBoxModels userItemBox.UserItemBoxes) userItemBox.UserItemBoxes {
	existingItemMaps := userItemBoxModels.SetExistingItemMaps()
	updateUserItemBoxModels := userItemBox.NewUserItemBoxes()

	for _, itemModel := range items {
		if existingItemMaps[itemModel.MasterItemId] {
			for _, userItemBoxModel := range userItemBoxModels {
				if userItemBoxModel.MasterItemId == itemModel.MasterItemId {
					updateUserItemBoxModels = append(updateUserItemBoxModels, userItemBox.SetUserItemBox(userItemBoxModel.UserId, userItemBoxModel.MasterItemId, userItemBoxModel.Count+itemModel.Count))
					break
				}
			}
		} else {
			updateUserItemBoxModels = append(updateUserItemBoxModels, userItemBox.SetUserItemBox(userId, itemModel.MasterItemId, itemModel.Count))
		}
	}

	return updateUserItemBoxModels
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
