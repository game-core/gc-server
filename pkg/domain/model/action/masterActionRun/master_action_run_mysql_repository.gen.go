// Package masterActionRun 実行されるアクション
//
//go:generate mockgen -source=./master_action_run_mysql_repository.gen.go -destination=./master_action_run_mysql_repository_mock.gen.go -package=masterActionRun
package masterActionRun

import (
	"context"

	"gorm.io/gorm"
)

type MasterActionRunMysqlRepository interface {
	Find(ctx context.Context, masterActionRunId int64) (*MasterActionRun, error)
	FindOrNil(ctx context.Context, masterActionRunId int64) (*MasterActionRun, error)
	FindByMasterActionId(ctx context.Context, masterActionId int64) (*MasterActionRun, error)
	FindOrNilByMasterActionId(ctx context.Context, masterActionId int64) (*MasterActionRun, error)
	FindList(ctx context.Context) (MasterActionRuns, error)
	FindListByMasterActionId(ctx context.Context, masterActionId int64) (MasterActionRuns, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterActionRuns) (MasterActionRuns, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterActionRun) (*MasterActionRun, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterActionRun) error
}
