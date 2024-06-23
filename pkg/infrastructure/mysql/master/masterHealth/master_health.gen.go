// Package masterHealth ヘルスチェック
package masterHealth

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type MasterHealths []*MasterHealth

type MasterHealth struct {
	HealthId         int64
	Name             string
	MasterHealthEnum masterHealth.MasterHealthEnum
}

func NewMasterHealth() *MasterHealth {
	return &MasterHealth{}
}

func NewMasterHealths() MasterHealths {
	return MasterHealths{}
}

func SetMasterHealth(healthId int64, name string, masterHealthEnum masterHealth.MasterHealthEnum) *MasterHealth {
	return &MasterHealth{
		HealthId:         healthId,
		Name:             name,
		MasterHealthEnum: masterHealthEnum,
	}
}

func (t *MasterHealth) TableName() string {
	return "master_health"
}
