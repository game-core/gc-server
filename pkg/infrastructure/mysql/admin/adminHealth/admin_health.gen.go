// Package adminHealth ヘルスチェック
package adminHealth

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/adminHealth"
)

type AdminHealths []*AdminHealth

type AdminHealth struct {
	HealthId        int64
	Name            string
	AdminHealthEnum adminHealth.AdminHealthEnum
}

func NewAdminHealth() *AdminHealth {
	return &AdminHealth{}
}

func NewAdminHealths() AdminHealths {
	return AdminHealths{}
}

func SetAdminHealth(healthId int64, name string, adminHealthEnum adminHealth.AdminHealthEnum) *AdminHealth {
	return &AdminHealth{
		HealthId:        healthId,
		Name:            name,
		AdminHealthEnum: adminHealthEnum,
	}
}

func (t *AdminHealth) TableName() string {
	return "admin_health"
}
