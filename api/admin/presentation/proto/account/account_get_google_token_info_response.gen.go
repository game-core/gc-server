// Package account アカウントGoogleToken情報取得レスポンス
package account

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleTokenInfo"
)

type AccountGetGoogleTokenInfoResponses []*AccountGetGoogleTokenInfoResponse

func NewAccountGetGoogleTokenInfoResponse() *AccountGetGoogleTokenInfoResponse {
	return &AccountGetGoogleTokenInfoResponse{}
}

func NewAccountGetGoogleTokenInfoResponses() AccountGetGoogleTokenInfoResponses {
	return AccountGetGoogleTokenInfoResponses{}
}

func SetAccountGetGoogleTokenInfoResponse(adminAccountGoogleTokenInfo *adminAccountGoogleTokenInfo.AdminAccountGoogleTokenInfo) *AccountGetGoogleTokenInfoResponse {
	return &AccountGetGoogleTokenInfoResponse{
		AdminAccountGoogleTokenInfo: adminAccountGoogleTokenInfo,
	}
}
