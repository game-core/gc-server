// Package masterLoginBonus ログインボーナス
//
//go:generate mockgen -source=./master_login_bonus_mysql_repository.gen.go -destination=./master_login_bonus_mysql_repository_mock.gen.go -package=masterLoginBonus
package masterLoginBonus

import (
	"context"

	"gorm.io/gorm"
)

type MasterLoginBonusMysqlRepository interface {
	Find(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonus, error)
	FindOrNil(ctx context.Context, masterLoginBonusId int64) (*MasterLoginBonus, error)
	FindByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error)
	FindOrNilByMasterEventId(ctx context.Context, masterEventId int64) (*MasterLoginBonus, error)
	FindList(ctx context.Context) (MasterLoginBonuses, error)
	FindListByMasterEventId(ctx context.Context, masterEventId int64) (MasterLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) (*MasterLoginBonus, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) (MasterLoginBonuses, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterLoginBonus) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterLoginBonuses) error
}
