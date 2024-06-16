// Package userExchange ユーザー交換
package userExchange

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

func SetUserExchange(userId string, masterExchangeId int64, resetAt *timestamppb.Timestamp) *UserExchange {
	return &UserExchange{
		UserId:           userId,
		MasterExchangeId: masterExchangeId,
		ResetAt:          resetAt,
	}
}
