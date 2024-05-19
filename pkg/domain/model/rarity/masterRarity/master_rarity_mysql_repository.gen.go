// Package masterRarity レアリティ
//
//go:generate mockgen -source=./master_rarity_mysql_repository.gen.go -destination=./master_rarity_mysql_repository_mock.gen.go -package=masterRarity
package masterRarity

import (
	"context"

	"gorm.io/gorm"
)

type MasterRarityMysqlRepository interface {
	Find(ctx context.Context, masterRarityId int64) (*MasterRarity, error)
	FindOrNil(ctx context.Context, masterRarityId int64) (*MasterRarity, error)
	FindByMasterRarityEnum(ctx context.Context, masterRarityEnum MasterRarityEnum) (*MasterRarity, error)
	FindOrNilByMasterRarityEnum(ctx context.Context, masterRarityEnum MasterRarityEnum) (*MasterRarity, error)
	FindList(ctx context.Context) (MasterRarities, error)
	FindListByMasterRarityEnum(ctx context.Context, masterRarityEnum MasterRarityEnum) (MasterRarities, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterRarity) (*MasterRarity, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterRarities) (MasterRarities, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterRarity) (*MasterRarity, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterRarity) error
}
