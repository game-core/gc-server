// Package account アカウントトークン取得レスポンス
package account

import (
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccountToken"
)

type AccountGetTokenResponses []*AccountGetTokenResponse

type AccountGetTokenResponse struct {
	UserAccountToken *userAccountToken.UserAccountToken
}

func NewAccountGetTokenResponse() *AccountGetTokenResponse {
	return &AccountGetTokenResponse{}
}

func NewAccountGetTokenResponses() AccountGetTokenResponses {
	return AccountGetTokenResponses{}
}

func SetAccountGetTokenResponse(userAccountToken *userAccountToken.UserAccountToken) *AccountGetTokenResponse {
	return &AccountGetTokenResponse{
		UserAccountToken: userAccountToken,
	}
}
