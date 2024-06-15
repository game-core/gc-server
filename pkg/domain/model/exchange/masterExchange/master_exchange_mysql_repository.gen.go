// Package masterExchange 交換
//
//go:generate mockgen -source=./master_exchange_mysql_repository.gen.go -destination=./master_exchange_mysql_repository_mock.gen.go -package=masterExchange
package masterExchange

import (
	"context"

	"gorm.io/gorm"
)

type MasterExchangeMysqlRepository interface {
	Find(ctx context.Context, masterExchangeId int64) (*MasterExchange, error)
	FindOrNil(ctx context.Context, masterExchangeId int64) (*MasterExchange, error)
	FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterExchange, error)
	FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterExchange, error)
	FindList(ctx context.Context) (MasterExchanges, error)
	FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterExchanges, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterExchange) (*MasterExchange, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) (MasterExchanges, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterExchange) (*MasterExchange, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) (MasterExchanges, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterExchange) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterExchanges) error
}
