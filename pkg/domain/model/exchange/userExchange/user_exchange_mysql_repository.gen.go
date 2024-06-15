// Package userExchange ユーザー交換
//
//go:generate mockgen -source=./user_exchange_mysql_repository.gen.go -destination=./user_exchange_mysql_repository_mock.gen.go -package=userExchange
package userExchange

import (
	"context"

	"gorm.io/gorm"
)

type UserExchangeMysqlRepository interface {
	Find(ctx context.Context, userId string, masterExchangeId int64) (*UserExchange, error)
	FindOrNil(ctx context.Context, userId string, masterExchangeId int64) (*UserExchange, error)
	FindList(ctx context.Context, userId string) (UserExchanges, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserExchange) (*UserExchange, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserExchanges) (UserExchanges, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserExchange) (*UserExchange, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms UserExchanges) (UserExchanges, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserExchange) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms UserExchanges) error
}
