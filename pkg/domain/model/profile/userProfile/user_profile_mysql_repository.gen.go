// Package userProfile ユーザープロフィール
//
//go:generate mockgen -source=./user_profile_mysql_repository.gen.go -destination=./user_profile_mysql_repository_mock.gen.go -package=userProfile
package userProfile

import (
	"context"

	"gorm.io/gorm"
)

type UserProfileMysqlRepository interface {
	Find(ctx context.Context, userId string) (*UserProfile, error)
	FindOrNil(ctx context.Context, userId string) (*UserProfile, error)
	FindList(ctx context.Context, userId string) (UserProfiles, error)
	Create(ctx context.Context, tx *gorm.DB, m *UserProfile) (*UserProfile, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms UserProfiles) (UserProfiles, error)
	Update(ctx context.Context, tx *gorm.DB, m *UserProfile) (*UserProfile, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms UserProfiles) (UserProfiles, error)
	Delete(ctx context.Context, tx *gorm.DB, m *UserProfile) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms UserProfiles) error
}
