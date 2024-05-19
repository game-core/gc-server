// Package account アカウント取得レスポンス
package account

import (
	"github.com/game-core/gc-server/api/game/presentation/server/account/userAccount"
)

func SetAccountGetResponse(userAccount *userAccount.UserAccount) *AccountGetResponse {
	return &AccountGetResponse{
		UserAccount: userAccount,
	}
}
