// Package account アカウント作成リクエスト
package account

type AccountCreateRequests []*AccountCreateRequest

func NewAccountCreateRequest() *AccountCreateRequest {
	return &AccountCreateRequest{}
}

func NewAccountCreateRequests() AccountCreateRequests {
	return AccountCreateRequests{}
}

func SetAccountCreateRequest(name string, password string) *AccountCreateRequest {
	return &AccountCreateRequest{
		Name:     name,
		Password: password,
	}
}
