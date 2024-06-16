// Package exchange 交換受け取りレスポンス
package exchange

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchangeItem"
)

func SetExchangeReceiveResponse(userExchange *userExchange.UserExchange, userExchangeItem *userExchangeItem.UserExchangeItem) *ExchangeReceiveResponse {
	return &ExchangeReceiveResponse{
		UserExchange:     userExchange,
		UserExchangeItem: userExchangeItem,
	}
}
