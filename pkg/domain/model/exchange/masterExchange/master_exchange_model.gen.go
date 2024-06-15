// Package masterExchange 交換
package masterExchange

type MasterExchanges []*MasterExchange

type MasterExchange struct {
	MasterExchangeId int64
	MasterEventId    int64
	Name             string
}

func NewMasterExchange() *MasterExchange {
	return &MasterExchange{}
}

func NewMasterExchanges() MasterExchanges {
	return MasterExchanges{}
}

func SetMasterExchange(masterExchangeId int64, masterEventId int64, name string) *MasterExchange {
	return &MasterExchange{
		MasterExchangeId: masterExchangeId,
		MasterEventId:    masterEventId,
		Name:             name,
	}
}
