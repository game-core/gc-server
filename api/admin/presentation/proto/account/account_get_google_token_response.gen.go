// Package account アカウントGoogleToken取得レスポンス
package account

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
)

type AccountGetGoogleTokenResponses []*AccountGetGoogleTokenResponse

func NewAccountGetGoogleTokenResponse() *AccountGetGoogleTokenResponse {
	return &AccountGetGoogleTokenResponse{}
}

func NewAccountGetGoogleTokenResponses() AccountGetGoogleTokenResponses {
	return AccountGetGoogleTokenResponses{}
}

func SetAccountGetGoogleTokenResponse(adminAccountGoogleToken *adminAccountGoogleToken.AdminAccountGoogleToken) *AccountGetGoogleTokenResponse {
	return &AccountGetGoogleTokenResponse{
		AdminAccountGoogleToken: adminAccountGoogleToken,
	}
}
