// Package account アカウントトークン取得リクエスト
package account

type AccountGetTokenRequests []*AccountGetTokenRequest

type AccountGetTokenRequest struct {
	UserId string
}

func NewAccountGetTokenRequest() *AccountGetTokenRequest {
	return &AccountGetTokenRequest{}
}

func NewAccountGetTokenRequests() AccountGetTokenRequests {
	return AccountGetTokenRequests{}
}

func SetAccountGetTokenRequest(userId string) *AccountGetTokenRequest {
	return &AccountGetTokenRequest{
		UserId: userId,
	}
}
