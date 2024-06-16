// Package exchange 交換更新リクエスト
package exchange

func SetExchangeUpdateRequest(userId string, masterExchangeId int64) *ExchangeUpdateRequest {
	return &ExchangeUpdateRequest{
		UserId:           userId,
		MasterExchangeId: masterExchangeId,
	}
}
