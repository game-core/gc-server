// Package exchange 交換更新レスポンス
package exchange

import (
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItems"
)

type ExchangeUpdateResponses []*ExchangeUpdateResponse

type ExchangeUpdateResponse struct {
	UserExchange      *userExchange.UserExchange
	UserExchangeItems userExchangeItems.UserExchangeItems
}

func NewExchangeUpdateResponse() *ExchangeUpdateResponse {
	return &ExchangeUpdateResponse{}
}

func NewExchangeUpdateResponses() ExchangeUpdateResponses {
	return ExchangeUpdateResponses{}
}

func SetExchangeUpdateResponse(userExchange *userExchange.UserExchange, userExchangeItems userExchangeItems.UserExchangeItems) *ExchangeUpdateResponse {
	return &ExchangeUpdateResponse{
		UserExchange:      userExchange,
		UserExchangeItems: userExchangeItems,
	}
}
