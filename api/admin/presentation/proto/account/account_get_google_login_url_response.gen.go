// Package account アカウントGoogleログインURL取得レスポンス
package account

type AccountGetGoogleLoginUrlResponses []*AccountGetGoogleLoginUrlResponse

func NewAccountGetGoogleLoginUrlResponse() *AccountGetGoogleLoginUrlResponse {
	return &AccountGetGoogleLoginUrlResponse{}
}

func NewAccountGetGoogleLoginUrlResponses() AccountGetGoogleLoginUrlResponses {
	return AccountGetGoogleLoginUrlResponses{}
}

func SetAccountGetGoogleLoginUrlResponse(url string) *AccountGetGoogleLoginUrlResponse {
	return &AccountGetGoogleLoginUrlResponse{
		Url: url,
	}
}
