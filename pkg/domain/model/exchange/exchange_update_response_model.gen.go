// Package exchange 交換更新レスポンス
package exchange

import (
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
)

type ExchangeUpdateResponses []*ExchangeUpdateResponse

type ExchangeUpdateResponse struct {
	UserExchange      *userExchange.UserExchange
	UserExchangeItems userExchangeItem.UserExchangeItems
}

func NewExchangeUpdateResponse() *ExchangeUpdateResponse {
	return &ExchangeUpdateResponse{}
}

func NewExchangeUpdateResponses() ExchangeUpdateResponses {
	return ExchangeUpdateResponses{}
}

func SetExchangeUpdateResponse(userExchange *userExchange.UserExchange, userExchangeItems userExchangeItem.UserExchangeItems) *ExchangeUpdateResponse {
	return &ExchangeUpdateResponse{
		UserExchange:      userExchange,
		UserExchangeItems: userExchangeItems,
	}
}
