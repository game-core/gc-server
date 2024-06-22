
// Package account アカウント作成レスポンス
package account

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
)

type AccountCreateResponses []*AccountCreateResponse

func NewAccountCreateResponse() *AccountCreateResponse {
			return &AccountCreateResponse{}
		}

		func NewAccountCreateResponses() AccountCreateResponses {
			return AccountCreateResponses{}
		}

		func SetAccountCreateResponse(userAccount *userAccount.UserAccount) *AccountCreateResponse {
			return &AccountCreateResponse{
				UserAccount: userAccount,
			}
		}
		
