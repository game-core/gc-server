// Package userExchangeItem ユーザー交換アイテム
package userExchangeItem

func SetUserExchangeItem(userId string, masterExchangeId int64, masterExchangeItemId int64, count int32) *UserExchangeItem {
	return &UserExchangeItem{
		UserId:               userId,
		MasterExchangeId:     masterExchangeId,
		MasterExchangeItemId: masterExchangeItemId,
		Count:                count,
	}
}
