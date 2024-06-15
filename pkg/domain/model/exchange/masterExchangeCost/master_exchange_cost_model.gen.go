// Package masterExchangeCost 交換コスト
package masterExchangeCost

type MasterExchangeCosts []*MasterExchangeCost

type MasterExchangeCost struct {
	MasterExchangeCostId int64
	MasterExchangeItemId int64
	Name                 string
	Count                int32
}

func NewMasterExchangeCost() *MasterExchangeCost {
	return &MasterExchangeCost{}
}

func NewMasterExchangeCosts() MasterExchangeCosts {
	return MasterExchangeCosts{}
}

func SetMasterExchangeCost(masterExchangeCostId int64, masterExchangeItemId int64, name string, count int32) *MasterExchangeCost {
	return &MasterExchangeCost{
		MasterExchangeCostId: masterExchangeCostId,
		MasterExchangeItemId: masterExchangeItemId,
		Name:                 name,
		Count:                count,
	}
}
