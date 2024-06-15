// Package userExchangeItem ユーザー交換アイテム
//
//go:generate mockgen -source=./user_exchange_item_mysql_repository.gen.go -destination=./user_exchange_item_mysql_repository_mock.gen.go -package=userExchangeItem
package userExchangeItem

import (
	"context"

	"gorm.io/gorm"
)

type UserExchangeItemMysqlRepository interface {
	Find(ctx context.Context, userId string, masterExchangeItemId int64) (*UserExchangeItem, error)
	FindOrNil(ctx context.Context, userId string, masterExchangeItemId int64) (*UserExchangeItem, error)
	FindByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (*UserExchangeItem, error)
	FindByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (*UserExchangeItem, error)
	FindOrNilByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (*UserExchangeItem, error)
	FindOrNilByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (*UserExchangeItem, error)
	FindList(ctx context.Context, userId string) (UserExchangeItems, error)
	FindListByUserIdAndMasterExchangeId(ctx context.Context, userId string, masterExchangeId int64) (UserExchangeItems, error)
	FindListByUserIdAndMasterExchangeItemId(ctx context.Context, userId string, masterExchangeItemId int64) (UserExchangeItems, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserExchangeItem) (*UserExchangeItem, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserExchangeItems) (UserExchangeItems, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserExchangeItem) (*UserExchangeItem, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms UserExchangeItems) (UserExchangeItems, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserExchangeItem) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms UserExchangeItems) error
}
