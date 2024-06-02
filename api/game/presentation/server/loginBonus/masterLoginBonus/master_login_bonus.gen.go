// Package masterLoginBonus ログインボーナス
package masterLoginBonus

func SetMasterLoginBonus(masterLoginBonusId int64, masterEventId int64, name string) *MasterLoginBonus {
	return &MasterLoginBonus{
		MasterLoginBonusId: masterLoginBonusId,
		MasterEventId:      masterEventId,
		Name:               name,
	}
}
