// Package account アカウントGoogleURL取得レスポンス
package account

import (
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleUrl"
)

type AccountGetGoogleUrlResponses []*AccountGetGoogleUrlResponse

type AccountGetGoogleUrlResponse struct {
	AdminAccountGoogleUrl *adminAccountGoogleUrl.AdminAccountGoogleUrl
}

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
