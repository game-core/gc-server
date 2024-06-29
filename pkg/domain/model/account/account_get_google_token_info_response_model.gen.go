// Package account アカウントGoogleToken情報取得レスポンス
package account

import (
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleTokenInfo"
)

type AccountGetGoogleTokenInfoResponses []*AccountGetGoogleTokenInfoResponse

type AccountGetGoogleTokenInfoResponse struct {
	AdminAccountGoogleTokenInfo *adminAccountGoogleTokenInfo.AdminAccountGoogleTokenInfo
}

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
