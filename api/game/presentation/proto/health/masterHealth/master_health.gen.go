// Package masterHealth ヘルスチェック
package masterHealth

func SetMasterHealth(healthId int64, name string, masterHealthType MasterHealthType) *MasterHealth {
	return &MasterHealth{
		HealthId:         healthId,
		Name:             name,
		MasterHealthType: masterHealthType,
	}
}
