// Package userExchangeItem ユーザー交換アイテム
package userExchangeItem

import (
	"time"
)

type UserExchangeItems []*UserExchangeItem

type UserExchangeItem struct {
	UserId               string
	MasterExchangeId     int64
	MasterExchangeItemId int64
	Count                int32
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func NewUserExchangeItem() *UserExchangeItem {
	return &UserExchangeItem{}
}

func NewUserExchangeItems() UserExchangeItems {
	return UserExchangeItems{}
}

func SetUserExchangeItem(userId string, masterExchangeId int64, masterExchangeItemId int64, count int32, createdAt time.Time, updatedAt time.Time) *UserExchangeItem {
	return &UserExchangeItem{
		UserId:               userId,
		MasterExchangeId:     masterExchangeId,
		MasterExchangeItemId: masterExchangeItemId,
		Count:                count,
		CreatedAt:            createdAt,
		UpdatedAt:            updatedAt,
	}
}

func (t *UserExchangeItem) TableName() string {
	return "user_exchange_item"
}
