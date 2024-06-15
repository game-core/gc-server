// Package userLoginBonus ユーザーログインボーナス
//
//go:generate mockgen -source=./user_login_bonus_mysql_repository.gen.go -destination=./user_login_bonus_mysql_repository_mock.gen.go -package=userLoginBonus
package userLoginBonus

import (
	"context"

	"gorm.io/gorm"
)

type UserLoginBonusMysqlRepository interface {
	Find(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindOrNil(ctx context.Context, userId string, masterLoginBonusId int64) (*UserLoginBonus, error)
	FindList(ctx context.Context, userId string) (UserLoginBonuses, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) (UserLoginBonuses, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) (*UserLoginBonus, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) (UserLoginBonuses, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserLoginBonus) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms UserLoginBonuses) error
}
