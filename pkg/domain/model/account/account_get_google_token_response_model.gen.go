// Package account アカウントGoogleToken取得レスポンス
package account

import (
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleToken"
)

type AccountGetGoogleTokenResponses []*AccountGetGoogleTokenResponse

type AccountGetGoogleTokenResponse struct {
	AdminAccountGoogleToken *adminAccountGoogleToken.AdminAccountGoogleToken
}

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
