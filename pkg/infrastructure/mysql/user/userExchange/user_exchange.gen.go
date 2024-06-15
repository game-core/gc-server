// Package userExchange ユーザー交換
package userExchange

import (
	"time"
)

type UserExchanges []*UserExchange

type UserExchange struct {
	UserId           string
	MasterExchangeId int64
	ResetAt          time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewUserExchange() *UserExchange {
	return &UserExchange{}
}

func NewUserExchanges() UserExchanges {
	return UserExchanges{}
}

func SetUserExchange(userId string, masterExchangeId int64, resetAt time.Time, createdAt time.Time, updatedAt time.Time) *UserExchange {
	return &UserExchange{
		UserId:           userId,
		MasterExchangeId: masterExchangeId,
		ResetAt:          resetAt,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}
}

func (t *UserExchange) TableName() string {
	return "user_exchange"
}
