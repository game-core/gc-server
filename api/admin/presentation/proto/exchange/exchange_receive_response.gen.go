
// Package exchange 交換受け取りレスポンス
package exchange

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchange"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchange"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchangeItem"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchangeItem"
)

type ExchangeReceiveResponses []*ExchangeReceiveResponse

func NewExchangeReceiveResponse() *ExchangeReceiveResponse {
			return &ExchangeReceiveResponse{}
		}

		func NewExchangeReceiveResponses() ExchangeReceiveResponses {
			return ExchangeReceiveResponses{}
		}

		func SetExchangeReceiveResponse(userExchange *userExchange.UserExchange,userExchangeItem *userExchangeItem.UserExchangeItem) *ExchangeReceiveResponse {
			return &ExchangeReceiveResponse{
				UserExchange: userExchange,
UserExchangeItem: userExchangeItem,
			}
		}
		
