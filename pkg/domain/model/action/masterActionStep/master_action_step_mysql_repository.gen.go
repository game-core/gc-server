// Package masterActionStep アクションステップ
//
//go:generate mockgen -source=./master_action_step_mysql_repository.gen.go -destination=./master_action_step_mysql_repository_mock.gen.go -package=masterActionStep
package masterActionStep

import (
	"context"

	"gorm.io/gorm"
)

type MasterActionStepMysqlRepository interface {
	Find(ctx context.Context, masterActionStepId int64) (*MasterActionStep, error)
	FindOrNil(ctx context.Context, masterActionStepId int64) (*MasterActionStep, error)
	FindByMasterActionStepEnum(ctx context.Context, masterActionStepEnum MasterActionStepEnum) (*MasterActionStep, error)
	FindOrNilByMasterActionStepEnum(ctx context.Context, masterActionStepEnum MasterActionStepEnum) (*MasterActionStep, error)
	FindList(ctx context.Context) (MasterActionSteps, error)
	FindListByMasterActionStepEnum(ctx context.Context, masterActionStepEnum MasterActionStepEnum) (MasterActionSteps, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionSteps) (MasterActionSteps, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionStep) (*MasterActionStep, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterActionSteps) (MasterActionSteps, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionStep) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterActionSteps) error
}
