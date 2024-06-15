// Package item アイテム消費レスポンス
package item

import (
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type ItemConsumeResponses []*ItemConsumeResponse

type ItemConsumeResponse struct {
	UserItemBoxes userItemBox.UserItemBoxes
}

func NewItemConsumeResponse() *ItemConsumeResponse {
	return &ItemConsumeResponse{}
}

func NewItemConsumeResponses() ItemConsumeResponses {
	return ItemConsumeResponses{}
}

func SetItemConsumeResponse(userItemBoxes userItemBox.UserItemBoxes) *ItemConsumeResponse {
	return &ItemConsumeResponse{
		UserItemBoxes: userItemBoxes,
	}
}
