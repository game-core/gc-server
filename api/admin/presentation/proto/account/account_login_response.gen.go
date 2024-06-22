
// Package account アカウントログインレスポンス
package account

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
)

type AccountLoginResponses []*AccountLoginResponse

func NewAccountLoginResponse() *AccountLoginResponse {
			return &AccountLoginResponse{}
		}

		func NewAccountLoginResponses() AccountLoginResponses {
			return AccountLoginResponses{}
		}

		func SetAccountLoginResponse(token string,userAccount *userAccount.UserAccount) *AccountLoginResponse {
			return &AccountLoginResponse{
				Token: token,
UserAccount: userAccount,
			}
		}
		
