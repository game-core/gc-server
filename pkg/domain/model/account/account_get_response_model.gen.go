// Package account アカウント取得レスポンス
package account

import (
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
)

type AccountGetResponses []*AccountGetResponse

type AccountGetResponse struct {
	UserAccount *userAccount.UserAccount
}

func NewAccountGetResponse() *AccountGetResponse {
	return &AccountGetResponse{}
}

func NewAccountGetResponses() AccountGetResponses {
	return AccountGetResponses{}
}

func SetAccountGetResponse(userAccount *userAccount.UserAccount) *AccountGetResponse {
	return &AccountGetResponse{
		UserAccount: userAccount,
	}
}
