
// Package account アカウント取得リクエスト
package account

import (
	
)

type AccountGetRequests []*AccountGetRequest

func NewAccountGetRequest() *AccountGetRequest {
			return &AccountGetRequest{}
		}

		func NewAccountGetRequests() AccountGetRequests {
			return AccountGetRequests{}
		}

		func SetAccountGetRequest(userId string) *AccountGetRequest {
			return &AccountGetRequest{
				UserId: userId,
			}
		}
		
