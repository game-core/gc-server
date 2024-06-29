// Package account アカウントGoogleTokenリフレッシュレスポンス
package account

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
)

type AccountRefreshGoogleTokenResponses []*AccountRefreshGoogleTokenResponse

func NewAccountRefreshGoogleTokenResponse() *AccountRefreshGoogleTokenResponse {
	return &AccountRefreshGoogleTokenResponse{}
}

func NewAccountRefreshGoogleTokenResponses() AccountRefreshGoogleTokenResponses {
	return AccountRefreshGoogleTokenResponses{}
}

func SetAccountRefreshGoogleTokenResponse(adminAccountGoogleToken *adminAccountGoogleToken.AdminAccountGoogleToken) *AccountRefreshGoogleTokenResponse {
	return &AccountRefreshGoogleTokenResponse{
		AdminAccountGoogleToken: adminAccountGoogleToken,
	}
}
