
// Package account アカウント取得レスポンス
package account

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
"github.com/game-core/gc-server/api/admin/presentation/proto/account/userAccount"
)

type AccountGetResponses []*AccountGetResponse

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
		
