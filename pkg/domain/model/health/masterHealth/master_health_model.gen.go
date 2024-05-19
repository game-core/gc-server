// Package masterHealth ヘルスチェック
package masterHealth

type MasterHealths []*MasterHealth

type MasterHealth struct {
	HealthId         int64
	Name             string
	MasterHealthType MasterHealthType
}

func NewMasterHealth() *MasterHealth {
	return &MasterHealth{}
}

func NewMasterHealths() MasterHealths {
	return MasterHealths{}
}

func SetMasterHealth(healthId int64, name string, masterHealthType MasterHealthType) *MasterHealth {
	return &MasterHealth{
		HealthId:         healthId,
		Name:             name,
		MasterHealthType: masterHealthType,
	}
}
