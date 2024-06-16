// Package exchange 交換更新レスポンス
package exchange

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchangeItem"
)

func SetExchangeUpdateResponse(userExchange *userExchange.UserExchange, userExchangeItems []*userExchangeItem.UserExchangeItem) *ExchangeUpdateResponse {
	return &ExchangeUpdateResponse{
		UserExchange:      userExchange,
		UserExchangeItems: userExchangeItems,
	}
}
