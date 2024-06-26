// Package account アカウントGoogleURL取得レスポンス
package account

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleUrl"
)

type AccountGetGoogleUrlResponses []*AccountGetGoogleUrlResponse

func NewAccountGetGoogleUrlResponse() *AccountGetGoogleUrlResponse {
	return &AccountGetGoogleUrlResponse{}
}

func NewAccountGetGoogleUrlResponses() AccountGetGoogleUrlResponses {
	return AccountGetGoogleUrlResponses{}
}

func SetAccountGetGoogleUrlResponse(adminAccountGoogleUrl *adminAccountGoogleUrl.AdminAccountGoogleUrl) *AccountGetGoogleUrlResponse {
	return &AccountGetGoogleUrlResponse{
		AdminAccountGoogleUrl: adminAccountGoogleUrl,
	}
}
