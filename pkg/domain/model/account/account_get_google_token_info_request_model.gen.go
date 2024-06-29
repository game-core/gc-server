// Package account アカウントGoogleToken情報取得リクエスト
package account

type AccountGetGoogleTokenInfoRequests []*AccountGetGoogleTokenInfoRequest

type AccountGetGoogleTokenInfoRequest struct {
	AccessToken string
}

func NewAccountGetGoogleTokenInfoRequest() *AccountGetGoogleTokenInfoRequest {
	return &AccountGetGoogleTokenInfoRequest{}
}

func NewAccountGetGoogleTokenInfoRequests() AccountGetGoogleTokenInfoRequests {
	return AccountGetGoogleTokenInfoRequests{}
}

func SetAccountGetGoogleTokenInfoRequest(accessToken string) *AccountGetGoogleTokenInfoRequest {
	return &AccountGetGoogleTokenInfoRequest{
		AccessToken: accessToken,
	}
}
