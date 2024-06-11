// Package masterLoginBonusItem ログインボーナスアイテム
package masterLoginBonusItem

func SetMasterLoginBonusItem(masterLoginBonusItemId int64, masterLoginBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterLoginBonusItem {
	return &MasterLoginBonusItem{
		MasterLoginBonusItemId:     masterLoginBonusItemId,
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterItemId:               masterItemId,
		Name:                       name,
		Count:                      count,
	}
}
