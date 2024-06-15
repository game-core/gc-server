// Package masterExchangeItem 交換アイテム
//
//go:generate mockgen -source=./master_exchange_item_mysql_repository.gen.go -destination=./master_exchange_item_mysql_repository_mock.gen.go -package=masterExchangeItem
package masterExchangeItem

import (
	"context"

	"gorm.io/gorm"
)

type MasterExchangeItemMysqlRepository interface {
	Find(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeItem, error)
	FindOrNil(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeItem, error)
	FindByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeItem, error)
	FindByMasterExchangeId(ctx context.Context, masterExchangeId int64) (*MasterExchangeItem, error)
	FindOrNilByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeItem, error)
	FindOrNilByMasterExchangeId(ctx context.Context, masterExchangeId int64) (*MasterExchangeItem, error)
	FindList(ctx context.Context) (MasterExchangeItems, error)
	FindListByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (MasterExchangeItems, error)
	FindListByMasterExchangeId(ctx context.Context, masterExchangeId int64) (MasterExchangeItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterExchangeItem) (*MasterExchangeItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterExchangeItems) (MasterExchangeItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterExchangeItem) (*MasterExchangeItem, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterExchangeItems) (MasterExchangeItems, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterExchangeItem) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterExchangeItems) error
}
