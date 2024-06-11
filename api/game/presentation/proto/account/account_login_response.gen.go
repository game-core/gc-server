// Package account アカウントログインレスポンス
package account

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/account/userAccount"
)

func SetAccountLoginResponse(token string, userAccount *userAccount.UserAccount) *AccountLoginResponse {
	return &AccountLoginResponse{
		Token:       token,
		UserAccount: userAccount,
	}
}
