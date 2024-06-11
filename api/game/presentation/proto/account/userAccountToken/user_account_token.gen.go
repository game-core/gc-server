// Package userAccountToken ユーザーアカウントトークン
package userAccountToken

func SetUserAccountToken(userId string, token string) *UserAccountToken {
	return &UserAccountToken{
		UserId: userId,
		Token:  token,
	}
}
