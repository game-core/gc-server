// Package masterHealth ヘルスチェック
package masterHealth

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type MasterHealths []*MasterHealth

type MasterHealth struct {
	HealthId         int64
	Name             string
	MasterHealthType masterHealth.MasterHealthType
}

func NewMasterHealth() *MasterHealth {
	return &MasterHealth{}
}

func NewMasterHealths() MasterHealths {
	return MasterHealths{}
}

func SetMasterHealth(healthId int64, name string, masterHealthType masterHealth.MasterHealthType) *MasterHealth {
	return &MasterHealth{
		HealthId:         healthId,
		Name:             name,
		MasterHealthType: masterHealthType,
	}
}

func (t *MasterHealth) TableName() string {
	return "master_health"
}
