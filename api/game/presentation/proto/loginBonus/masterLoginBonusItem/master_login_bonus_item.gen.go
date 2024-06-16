// Package masterLoginBonusItem ログインボーナスアイテム
package masterLoginBonusItem

type MasterLoginBonusItems []*MasterLoginBonusItem

func NewMasterLoginBonusItem() *MasterLoginBonusItem {
	return &MasterLoginBonusItem{}
}

func NewMasterLoginBonusItems() MasterLoginBonusItems {
	return MasterLoginBonusItems{}
}

func SetMasterLoginBonusItem(masterLoginBonusItemId int64, masterLoginBonusScheduleId int64, masterItemId int64, name string, count int32) *MasterLoginBonusItem {
	return &MasterLoginBonusItem{
		MasterLoginBonusItemId:     masterLoginBonusItemId,
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterItemId:               masterItemId,
		Name:                       name,
		Count:                      count,
	}
}
