// Package masterExchangeCost 交換コスト
//
//go:generate mockgen -source=./master_exchange_cost_mysql_repository.gen.go -destination=./master_exchange_cost_mysql_repository_mock.gen.go -package=masterExchangeCost
package masterExchangeCost

import (
	"context"

	"gorm.io/gorm"
)

type MasterExchangeCostMysqlRepository interface {
	Find(ctx context.Context, masterExchangeCostId int64) (*MasterExchangeCost, error)
	FindOrNil(ctx context.Context, masterExchangeCostId int64) (*MasterExchangeCost, error)
	FindByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeCost, error)
	FindOrNilByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (*MasterExchangeCost, error)
	FindList(ctx context.Context) (MasterExchangeCosts, error)
	FindListByMasterExchangeItemId(ctx context.Context, masterExchangeItemId int64) (MasterExchangeCosts, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterExchangeCost) (*MasterExchangeCost, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterExchangeCosts) (MasterExchangeCosts, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterExchangeCost) (*MasterExchangeCost, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterExchangeCosts) (MasterExchangeCosts, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterExchangeCost) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterExchangeCosts) error
}
