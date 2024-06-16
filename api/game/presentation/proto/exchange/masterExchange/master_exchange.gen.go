// Package masterExchange 交換
package masterExchange

func SetMasterExchange(masterExchangeId int64, masterEventId int64, name string) *MasterExchange {
	return &MasterExchange{
		MasterExchangeId: masterExchangeId,
		MasterEventId:    masterEventId,
		Name:             name,
	}
}
