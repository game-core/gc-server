// Package adminHealth ヘルスチェック
package adminHealth

type AdminHealths []*AdminHealth

type AdminHealth struct {
	HealthId        int64
	Name            string
	AdminHealthType AdminHealthType
}

func NewAdminHealth() *AdminHealth {
	return &AdminHealth{}
}

func NewAdminHealths() AdminHealths {
	return AdminHealths{}
}

func SetAdminHealth(healthId int64, name string, adminHealthType AdminHealthType) *AdminHealth {
	return &AdminHealth{
		HealthId:        healthId,
		Name:            name,
		AdminHealthType: adminHealthType,
	}
}
