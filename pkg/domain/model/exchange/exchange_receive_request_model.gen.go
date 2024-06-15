// Package exchange 交換受け取りリクエスト
package exchange

type ExchangeReceiveRequests []*ExchangeReceiveRequest

type ExchangeReceiveRequest struct {
	UserId               string
	MasterExchangeItemId int64
	Count                int32
}

func NewExchangeReceiveRequest() *ExchangeReceiveRequest {
	return &ExchangeReceiveRequest{}
}

func NewExchangeReceiveRequests() ExchangeReceiveRequests {
	return ExchangeReceiveRequests{}
}

func SetExchangeReceiveRequest(userId string, masterExchangeItemId int64, count int32) *ExchangeReceiveRequest {
	return &ExchangeReceiveRequest{
		UserId:               userId,
		MasterExchangeItemId: masterExchangeItemId,
		Count:                count,
	}
}
