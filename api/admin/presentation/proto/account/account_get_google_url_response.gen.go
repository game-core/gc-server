// Package account アカウントGoogleURL取得レスポンス
package account

type AccountGetGoogleUrlResponses []*AccountGetGoogleUrlResponse

func NewAccountGetGoogleUrlResponse() *AccountGetGoogleUrlResponse {
	return &AccountGetGoogleUrlResponse{}
}

func NewAccountGetGoogleUrlResponses() AccountGetGoogleUrlResponses {
	return AccountGetGoogleUrlResponses{}
}

func SetAccountGetGoogleUrlResponse(url string) *AccountGetGoogleUrlResponse {
	return &AccountGetGoogleUrlResponse{
		Url: url,
	}
}
