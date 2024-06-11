// Package account アカウントログインリクエスト
package account

func SetAccountLoginRequest(userId string, password string) *AccountLoginRequest {
	return &AccountLoginRequest{
		UserId:   userId,
		Password: password,
	}
}
