
// Package masterExchange 交換
package masterExchange

import (
	
)

type MasterExchanges []*MasterExchange

func NewMasterExchange() *MasterExchange {
			return &MasterExchange{}
		}

		func NewMasterExchanges() MasterExchanges {
			return MasterExchanges{}
		}

		func SetMasterExchange(masterExchangeId int64,masterEventId int64,name string) *MasterExchange {
			return &MasterExchange{
				MasterExchangeId: masterExchangeId,
MasterEventId: masterEventId,
Name: name,
			}
		}
		
