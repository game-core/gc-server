// Package userExchangeItem ユーザー交換アイテム
package userExchangeItem

type UserExchangeItems []*UserExchangeItem

type UserExchangeItem struct {
	UserId               string
	MasterExchangeId     int64
	MasterExchangeItemId int64
	Count                int32
}

func NewUserExchangeItem() *UserExchangeItem {
	return &UserExchangeItem{}
}

func NewUserExchangeItems() UserExchangeItems {
	return UserExchangeItems{}
}

func SetUserExchangeItem(userId string, masterExchangeId int64, masterExchangeItemId int64, count int32) *UserExchangeItem {
	return &UserExchangeItem{
		UserId:               userId,
		MasterExchangeId:     masterExchangeId,
		MasterExchangeItemId: masterExchangeItemId,
		Count:                count,
	}
}
