// Package account アカウント取得リクエスト
package account

type AccountGetRequests []*AccountGetRequest

type AccountGetRequest struct {
	UserId string
}

func NewAccountGetRequest() *AccountGetRequest {
	return &AccountGetRequest{}
}

func NewAccountGetRequests() AccountGetRequests {
	return AccountGetRequests{}
}

func SetAccountGetRequest(userId string) *AccountGetRequest {
	return &AccountGetRequest{
		UserId: userId,
	}
}
