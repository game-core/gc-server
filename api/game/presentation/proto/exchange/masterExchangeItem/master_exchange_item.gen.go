// Package masterExchangeItem 交換アイテム
package masterExchangeItem

func SetMasterExchangeItem(masterExchangeItemId int64, masterExchangeId int64, masterItemId int64, name string, count int32) *MasterExchangeItem {
	return &MasterExchangeItem{
		MasterExchangeItemId: masterExchangeItemId,
		MasterExchangeId:     masterExchangeId,
		MasterItemId:         masterItemId,
		Name:                 name,
		Count:                count,
	}
}
