// Package account アカウントGoogleログインURL取得リクエスト
package account

type AccountGetGoogleLoginUrlRequests []*AccountGetGoogleLoginUrlRequest

func NewAccountGetGoogleLoginUrlRequest() *AccountGetGoogleLoginUrlRequest {
	return &AccountGetGoogleLoginUrlRequest{}
}

func NewAccountGetGoogleLoginUrlRequests() AccountGetGoogleLoginUrlRequests {
	return AccountGetGoogleLoginUrlRequests{}
}

func SetAccountGetGoogleLoginUrlRequest(message string) *AccountGetGoogleLoginUrlRequest {
	return &AccountGetGoogleLoginUrlRequest{
		Message: message,
	}
}
