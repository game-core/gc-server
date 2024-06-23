// Package account アカウントGoogleログインToken取得レスポンス
package account

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
)

type AccountGetGoogleLoginTokenResponses []*AccountGetGoogleLoginTokenResponse

func NewAccountGetGoogleLoginTokenResponse() *AccountGetGoogleLoginTokenResponse {
	return &AccountGetGoogleLoginTokenResponse{}
}

func NewAccountGetGoogleLoginTokenResponses() AccountGetGoogleLoginTokenResponses {
	return AccountGetGoogleLoginTokenResponses{}
}

func SetAccountGetGoogleLoginTokenResponse(adminAccountGoogleToken *adminAccountGoogleToken.AdminAccountGoogleToken) *AccountGetGoogleLoginTokenResponse {
	return &AccountGetGoogleLoginTokenResponse{
		AdminAccountGoogleToken: adminAccountGoogleToken,
	}
}
