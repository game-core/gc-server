// Package account アカウントGoogleTokenリフレッシュリクエスト
package account

type AccountRefreshGoogleTokenRequests []*AccountRefreshGoogleTokenRequest

func NewAccountRefreshGoogleTokenRequest() *AccountRefreshGoogleTokenRequest {
	return &AccountRefreshGoogleTokenRequest{}
}

func NewAccountRefreshGoogleTokenRequests() AccountRefreshGoogleTokenRequests {
	return AccountRefreshGoogleTokenRequests{}
}

func SetAccountRefreshGoogleTokenRequest(refreshToken string) *AccountRefreshGoogleTokenRequest {
	return &AccountRefreshGoogleTokenRequest{
		RefreshToken: refreshToken,
	}
}
