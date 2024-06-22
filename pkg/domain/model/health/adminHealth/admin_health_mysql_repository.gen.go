// Package adminHealth ヘルスチェック
//
//go:generate mockgen -source=./admin_health_mysql_repository.gen.go -destination=./admin_health_mysql_repository_mock.gen.go -package=adminHealth
package adminHealth

import (
	"context"

	"gorm.io/gorm"
)

type AdminHealthMysqlRepository interface {
	Find(ctx context.Context, healthId int64) (*AdminHealth, error)
	FindOrNil(ctx context.Context, healthId int64) (*AdminHealth, error)
	FindList(ctx context.Context) (AdminHealths, error)
	Create(ctx context.Context, tx *gorm.DB, m *AdminHealth) (*AdminHealth, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms AdminHealths) (AdminHealths, error)
	Update(ctx context.Context, tx *gorm.DB, m *AdminHealth) (*AdminHealth, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms AdminHealths) (AdminHealths, error)
	Delete(ctx context.Context, tx *gorm.DB, m *AdminHealth) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms AdminHealths) error
}
