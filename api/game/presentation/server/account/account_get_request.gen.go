// Package account アカウント取得リクエスト
package account

func SetAccountGetRequest(userId string) *AccountGetRequest {
	return &AccountGetRequest{
		UserId: userId,
	}
}
