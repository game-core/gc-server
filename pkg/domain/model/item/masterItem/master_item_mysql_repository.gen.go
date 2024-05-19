// Package masterItem アイテム
//
//go:generate mockgen -source=./master_item_mysql_repository.gen.go -destination=./master_item_mysql_repository_mock.gen.go -package=masterItem
package masterItem

import (
	"context"

	"gorm.io/gorm"
)

type MasterItemMysqlRepository interface {
	Find(ctx context.Context, masterItemId int64) (*MasterItem, error)
	FindOrNil(ctx context.Context, masterItemId int64) (*MasterItem, error)
	FindList(ctx context.Context) (MasterItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterItem) (*MasterItem, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterItems) (MasterItems, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterItem) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterItems) error
}
