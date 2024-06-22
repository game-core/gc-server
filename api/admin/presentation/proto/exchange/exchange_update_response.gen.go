
// Package exchange 交換更新レスポンス
package exchange

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchange"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchange"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchangeItem"
"github.com/game-core/gc-server/api/admin/presentation/proto/exchange/userExchangeItem"
)

type ExchangeUpdateResponses []*ExchangeUpdateResponse

func NewExchangeUpdateResponse() *ExchangeUpdateResponse {
			return &ExchangeUpdateResponse{}
		}

		func NewExchangeUpdateResponses() ExchangeUpdateResponses {
			return ExchangeUpdateResponses{}
		}

		func SetExchangeUpdateResponse(userExchange *userExchange.UserExchange,userExchangeItems []*userExchangeItem.UserExchangeItem) *ExchangeUpdateResponse {
			return &ExchangeUpdateResponse{
				UserExchange: userExchange,
UserExchangeItems: userExchangeItems,
			}
		}
		
