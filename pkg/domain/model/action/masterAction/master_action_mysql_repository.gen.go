// Package masterAction アクション
//
//go:generate mockgen -source=./master_action_mysql_repository.gen.go -destination=./master_action_mysql_repository_mock.gen.go -package=masterAction
package masterAction

import (
	"context"

	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	"gorm.io/gorm"
)

type MasterActionMysqlRepository interface {
	Find(ctx context.Context, masterActionId int64) (*MasterAction, error)
	FindOrNil(ctx context.Context, masterActionId int64) (*MasterAction, error)
	FindByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*MasterAction, error)
	FindByTargetId(ctx context.Context, targetId *int64) (*MasterAction, error)
	FindByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*MasterAction, error)
	FindOrNilByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (*MasterAction, error)
	FindOrNilByTargetId(ctx context.Context, targetId *int64) (*MasterAction, error)
	FindOrNilByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (*MasterAction, error)
	FindList(ctx context.Context) (MasterActions, error)
	FindListByMasterActionStepEnum(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum) (MasterActions, error)
	FindListByTargetId(ctx context.Context, targetId *int64) (MasterActions, error)
	FindListByMasterActionStepEnumAndTargetId(ctx context.Context, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) (MasterActions, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterAction) (*MasterAction, error)
	UpdateList(ctx context.Context, tx *gorm.DB, ms MasterActions) (MasterActions, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterAction) error
	DeleteList(ctx context.Context, tx *gorm.DB, ms MasterActions) error
}
