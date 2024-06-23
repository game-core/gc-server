// Package account アカウントGoogleログインToken取得リクエスト
package account

type AccountGetGoogleLoginTokenRequests []*AccountGetGoogleLoginTokenRequest

func NewAccountGetGoogleLoginTokenRequest() *AccountGetGoogleLoginTokenRequest {
	return &AccountGetGoogleLoginTokenRequest{}
}

func NewAccountGetGoogleLoginTokenRequests() AccountGetGoogleLoginTokenRequests {
	return AccountGetGoogleLoginTokenRequests{}
}

func SetAccountGetGoogleLoginTokenRequest(code string) *AccountGetGoogleLoginTokenRequest {
	return &AccountGetGoogleLoginTokenRequest{
		Code: code,
	}
}
