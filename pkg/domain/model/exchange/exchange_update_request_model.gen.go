// Package exchange 交換更新リクエスト
package exchange

type ExchangeUpdateRequests []*ExchangeUpdateRequest

type ExchangeUpdateRequest struct {
	UserId           string
	MasterExchangeId int64
}

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
