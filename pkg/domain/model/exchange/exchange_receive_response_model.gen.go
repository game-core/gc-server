// Package exchange 交換受け取りレスポンス
package exchange

import (
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchange"
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
)

type ExchangeReceiveResponses []*ExchangeReceiveResponse

type ExchangeReceiveResponse struct {
	UserExchange     *userExchange.UserExchange
	UserExchangeItem *userExchangeItem.UserExchangeItem
}

func NewExchangeReceiveResponse() *ExchangeReceiveResponse {
	return &ExchangeReceiveResponse{}
}

func NewExchangeReceiveResponses() ExchangeReceiveResponses {
	return ExchangeReceiveResponses{}
}

func SetExchangeReceiveResponse(userExchange *userExchange.UserExchange, userExchangeItem *userExchangeItem.UserExchangeItem) *ExchangeReceiveResponse {
	return &ExchangeReceiveResponse{
		UserExchange:     userExchange,
		UserExchangeItem: userExchangeItem,
	}
}
