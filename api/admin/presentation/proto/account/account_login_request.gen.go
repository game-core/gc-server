
// Package account アカウントログインリクエスト
package account

import (
	
)

type AccountLoginRequests []*AccountLoginRequest

func NewAccountLoginRequest() *AccountLoginRequest {
			return &AccountLoginRequest{}
		}

		func NewAccountLoginRequests() AccountLoginRequests {
			return AccountLoginRequests{}
		}

		func SetAccountLoginRequest(userId string,password string) *AccountLoginRequest {
			return &AccountLoginRequest{
				UserId: userId,
Password: password,
			}
		}
		
