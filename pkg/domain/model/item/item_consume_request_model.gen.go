// Package item アイテム消費リクエスト
package item

type ItemConsumeRequests []*ItemConsumeRequest

type ItemConsumeRequest struct {
	UserId string
	Items  Items
}

func NewItemConsumeRequest() *ItemConsumeRequest {
	return &ItemConsumeRequest{}
}

func NewItemConsumeRequests() ItemConsumeRequests {
	return ItemConsumeRequests{}
}

func SetItemConsumeRequest(userId string, items Items) *ItemConsumeRequest {
	return &ItemConsumeRequest{
		UserId: userId,
		Items:  items,
	}
}
