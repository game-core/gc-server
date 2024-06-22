// Package exchange 交換更新リクエスト
package exchange

type ExchangeUpdateRequests []*ExchangeUpdateRequest

func NewExchangeUpdateRequest() *ExchangeUpdateRequest {
	return &ExchangeUpdateRequest{}
}

func NewExchangeUpdateRequests() ExchangeUpdateRequests {
	return ExchangeUpdateRequests{}
}

func SetExchangeUpdateRequest(userId string, masterExchangeId int64) *ExchangeUpdateRequest {
	return &ExchangeUpdateRequest{
		UserId:           userId,
		MasterExchangeId: masterExchangeId,
	}
}
