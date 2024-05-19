// Package commonHealth ヘルスチェック
//
//go:generate mockgen -source=./common_health_mysql_repository.gen.go -destination=./common_health_mysql_repository_mock.gen.go -package=commonHealth
package commonHealth

import (
	"context"

	"gorm.io/gorm"
)

type CommonHealthMysqlRepository interface {
	Find(ctx context.Context, healthId int64) (*CommonHealth, error)
	FindOrNil(ctx context.Context, healthId int64) (*CommonHealth, error)
	FindList(ctx context.Context) (CommonHealths, error)
	Create(ctx context.Context, tx *gorm.DB, m *CommonHealth) (*CommonHealth, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms CommonHealths) (CommonHealths, error)
	Update(ctx context.Context, tx *gorm.DB, m *CommonHealth) (*CommonHealth, error)
	Delete(ctx context.Context, tx *gorm.DB, m *CommonHealth) error
}
