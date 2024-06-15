// Package masterExchangeItem 交換アイテム
package masterExchangeItem

type MasterExchangeItems []*MasterExchangeItem

type MasterExchangeItem struct {
	MasterExchangeItemId int64
	MasterExchangeId     int64
	MasterItemId         int64
	Name                 string
	Count                int32
}

func NewMasterExchangeItem() *MasterExchangeItem {
	return &MasterExchangeItem{}
}

func NewMasterExchangeItems() MasterExchangeItems {
	return MasterExchangeItems{}
}

func SetMasterExchangeItem(masterExchangeItemId int64, masterExchangeId int64, masterItemId int64, name string, count int32) *MasterExchangeItem {
	return &MasterExchangeItem{
		MasterExchangeItemId: masterExchangeItemId,
		MasterExchangeId:     masterExchangeId,
		MasterItemId:         masterItemId,
		Name:                 name,
		Count:                count,
	}
}
