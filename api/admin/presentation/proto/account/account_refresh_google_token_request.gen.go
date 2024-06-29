// Package account アカウントGoogleTokenリフレッシュリクエスト
package account

type AccountRefreshGoogleTokenRequests []*AccountRefreshGoogleTokenRequest

func NewAccountRefreshGoogleTokenRequest() *AccountRefreshGoogleTokenRequest {
	return &AccountRefreshGoogleTokenRequest{}
}

func NewAccountRefreshGoogleTokenRequests() AccountRefreshGoogleTokenRequests {
	return AccountRefreshGoogleTokenRequests{}
}

func SetAccountRefreshGoogleTokenRequest() *AccountRefreshGoogleTokenRequest {
	return &AccountRefreshGoogleTokenRequest{}
}
