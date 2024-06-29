// Package account アカウントGoogleURL取得リクエスト
package account

type AccountGetGoogleUrlRequests []*AccountGetGoogleUrlRequest

type AccountGetGoogleUrlRequest struct {
}

func NewAccountGetGoogleUrlRequest() *AccountGetGoogleUrlRequest {
	return &AccountGetGoogleUrlRequest{}
}

func NewAccountGetGoogleUrlRequests() AccountGetGoogleUrlRequests {
	return AccountGetGoogleUrlRequests{}
}

func SetAccountGetGoogleUrlRequest() *AccountGetGoogleUrlRequest {
	return &AccountGetGoogleUrlRequest{}
}
