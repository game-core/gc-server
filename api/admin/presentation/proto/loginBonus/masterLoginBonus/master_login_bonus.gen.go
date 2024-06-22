
// Package masterLoginBonus ログインボーナス
package masterLoginBonus

import (
	
)

type MasterLoginBonuses []*MasterLoginBonus

func NewMasterLoginBonus() *MasterLoginBonus {
			return &MasterLoginBonus{}
		}

		func NewMasterLoginBonuses() MasterLoginBonuses {
			return MasterLoginBonuses{}
		}

		func SetMasterLoginBonus(masterLoginBonusId int64,masterEventId int64,name string) *MasterLoginBonus {
			return &MasterLoginBonus{
				MasterLoginBonusId: masterLoginBonusId,
MasterEventId: masterEventId,
Name: name,
			}
		}
		
