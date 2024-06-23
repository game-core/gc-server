// Package adminHealth ヘルスチェック
package adminHealth

type AdminHealths []*AdminHealth

func NewAdminHealth() *AdminHealth {
	return &AdminHealth{}
}

func NewAdminHealths() AdminHealths {
	return AdminHealths{}
}

func SetAdminHealth(healthId int64, name string, adminHealthEnum AdminHealthEnum) *AdminHealth {
	return &AdminHealth{
		HealthId:        healthId,
		Name:            name,
		AdminHealthEnum: adminHealthEnum,
	}
}
