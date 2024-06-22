
// Package exchange 交換受け取りリクエスト
package exchange

import (
	
)

type ExchangeReceiveRequests []*ExchangeReceiveRequest

func NewExchangeReceiveRequest() *ExchangeReceiveRequest {
			return &ExchangeReceiveRequest{}
		}

		func NewExchangeReceiveRequests() ExchangeReceiveRequests {
			return ExchangeReceiveRequests{}
		}

		func SetExchangeReceiveRequest(userId string,masterExchangeItemId int64,count int32) *ExchangeReceiveRequest {
			return &ExchangeReceiveRequest{
				UserId: userId,
MasterExchangeItemId: masterExchangeItemId,
Count: count,
			}
		}
		
