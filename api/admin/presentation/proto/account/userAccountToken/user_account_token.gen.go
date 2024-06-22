
// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

import (
	
)

type UserAccountTokens []*UserAccountToken

func NewUserAccountToken() *UserAccountToken {
			return &UserAccountToken{}
		}

		func NewUserAccountTokens() UserAccountTokens {
			return UserAccountTokens{}
		}

		func SetUserAccountToken(userId string,token string) *UserAccountToken {
			return &UserAccountToken{
				UserId: userId,
Token: token,
			}
		}
		
