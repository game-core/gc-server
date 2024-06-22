// Package masterExchangeCost 交換コスト
package masterExchangeCost

type MasterExchangeCosts []*MasterExchangeCost

func NewMasterExchangeCost() *MasterExchangeCost {
	return &MasterExchangeCost{}
}

func NewMasterExchangeCosts() MasterExchangeCosts {
	return MasterExchangeCosts{}
}

func SetMasterExchangeCost(masterExchangeCostId int64, masterExchangeItemId int64, masterItemId int64, name string, count int32) *MasterExchangeCost {
	return &MasterExchangeCost{
		MasterExchangeCostId: masterExchangeCostId,
		MasterExchangeItemId: masterExchangeItemId,
		MasterItemId:         masterItemId,
		Name:                 name,
		Count:                count,
	}
}
