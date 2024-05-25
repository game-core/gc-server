// Package item アイテム受け取りレスポンス
package item

import (
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type ItemReceiveResponses []*ItemReceiveResponse

type ItemReceiveResponse struct {
	UserItemBoxes userItemBox.UserItemBoxes
}

func NewItemReceiveResponse() *ItemReceiveResponse {
	return &ItemReceiveResponse{}
}

func NewItemReceiveResponses() ItemReceiveResponses {
	return ItemReceiveResponses{}
}

func SetItemReceiveResponse(userItemBoxes userItemBox.UserItemBoxes) *ItemReceiveResponse {
	return &ItemReceiveResponse{
		UserItemBoxes: userItemBoxes,
	}
}
