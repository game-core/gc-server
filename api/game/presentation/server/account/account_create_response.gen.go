// Package account アカウント作成レスポンス
package account

import (
	"github.com/game-core/gc-server/api/game/presentation/server/account/userAccount"
)

func SetAccountCreateResponse(userAccount *userAccount.UserAccount) *AccountCreateResponse {
	return &AccountCreateResponse{
		UserAccount: userAccount,
	}
}
