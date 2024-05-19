// Package masterHealth ヘルスチェック
//
//go:generate mockgen -source=./master_health_mysql_repository.gen.go -destination=./master_health_mysql_repository_mock.gen.go -package=masterHealth
package masterHealth

import (
	"context"

	"gorm.io/gorm"
)

type MasterHealthMysqlRepository interface {
	Find(ctx context.Context, healthId int64) (*MasterHealth, error)
	FindOrNil(ctx context.Context, healthId int64) (*MasterHealth, error)
	FindList(ctx context.Context) (MasterHealths, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterHealth) (*MasterHealth, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterHealths) (MasterHealths, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterHealth) (*MasterHealth, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterHealth) error
}
