
// Package masterHealth ヘルスチェック
package masterHealth

import (
	
)

type MasterHealths []*MasterHealth

func NewMasterHealth() *MasterHealth {
			return &MasterHealth{}
		}

		func NewMasterHealths() MasterHealths {
			return MasterHealths{}
		}

		func SetMasterHealth(healthId int64,name string,masterHealthType MasterHealthType) *MasterHealth {
			return &MasterHealth{
				HealthId: healthId,
Name: name,
MasterHealthType: masterHealthType,
			}
		}
		
