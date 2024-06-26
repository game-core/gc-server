// Package userExchangeItem ユーザー交換アイテム
package userExchangeItem

import (
	"github.com/game-core/gc-server/pkg/domain/model/exchange/userExchangeItem"
)

func SetUserExchangeItems(userExchangeItemModels userExchangeItem.UserExchangeItems) UserExchangeItems {
	userExchangeItems := NewUserExchangeItems()
	for _, userExchangeItemModel := range userExchangeItemModels {
		userExchangeItems = append(
			userExchangeItems,
			SetUserExchangeItem(
				userExchangeItemModel.UserId,
				userExchangeItemModel.MasterExchangeId,
				userExchangeItemModel.MasterExchangeItemId,
				userExchangeItemModel.Count,
			),
		)
	}

	return userExchangeItems
}
