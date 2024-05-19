// Package masterResource リソース
//
//go:generate mockgen -source=./master_resource_mysql_repository.gen.go -destination=./master_resource_mysql_repository_mock.gen.go -package=masterResource
package masterResource

import (
	"context"

	"gorm.io/gorm"
)

type MasterResourceMysqlRepository interface {
	Find(ctx context.Context, masterResourceId int64) (*MasterResource, error)
	FindOrNil(ctx context.Context, masterResourceId int64) (*MasterResource, error)
	FindByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (*MasterResource, error)
	FindOrNilByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (*MasterResource, error)
	FindList(ctx context.Context) (MasterResources, error)
	FindListByMasterResourceEnum(ctx context.Context, masterResourceEnum MasterResourceEnum) (MasterResources, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterResources) (MasterResources, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterResource) (*MasterResource, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterResources) (MasterResources, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterResource) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterResources) error
}
