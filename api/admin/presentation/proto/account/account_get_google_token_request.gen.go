// Package account アカウントGoogleToken取得リクエスト
package account

type AccountGetGoogleTokenRequests []*AccountGetGoogleTokenRequest

func NewAccountGetGoogleTokenRequest() *AccountGetGoogleTokenRequest {
	return &AccountGetGoogleTokenRequest{}
}

func NewAccountGetGoogleTokenRequests() AccountGetGoogleTokenRequests {
	return AccountGetGoogleTokenRequests{}
}

func SetAccountGetGoogleTokenRequest(code string) *AccountGetGoogleTokenRequest {
	return &AccountGetGoogleTokenRequest{
		Code: code,
	}
}
