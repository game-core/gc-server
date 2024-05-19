// Package masterActionTrigger アクショントリガー
//
//go:generate mockgen -source=./master_action_trigger_mysql_repository.gen.go -destination=./master_action_trigger_mysql_repository_mock.gen.go -package=masterActionTrigger
package masterActionTrigger

import (
	"context"

	"gorm.io/gorm"
)

type MasterActionTriggerMysqlRepository interface {
	Find(ctx context.Context, masterActionTriggerId int64) (*MasterActionTrigger, error)
	FindOrNil(ctx context.Context, masterActionTriggerId int64) (*MasterActionTrigger, error)
	FindByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum MasterActionTriggerEnum) (*MasterActionTrigger, error)
	FindOrNilByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum MasterActionTriggerEnum) (*MasterActionTrigger, error)
	FindList(ctx context.Context) (MasterActionTriggers, error)
	FindListByMasterActionTriggerEnum(ctx context.Context, masterActionTriggerEnum MasterActionTriggerEnum) (MasterActionTriggers, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) (*MasterActionTrigger, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionTriggers) (MasterActionTriggers, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) (*MasterActionTrigger, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionTrigger) error
}
