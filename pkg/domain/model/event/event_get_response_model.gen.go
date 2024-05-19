// Package event イベント取得レスポンス
package event

import (
	"github.com/game-core/gc-server/pkg/domain/model/event/masterEvent"
)

type EventGetResponses []*EventGetResponse

type EventGetResponse struct {
	MasterEvent *masterEvent.MasterEvent
}

func NewEventGetResponse() *EventGetResponse {
	return &EventGetResponse{}
}

func NewEventGetResponses() EventGetResponses {
	return EventGetResponses{}
}

func SetEventGetResponse(masterEvent *masterEvent.MasterEvent) *EventGetResponse {
	return &EventGetResponse{
		MasterEvent: masterEvent,
	}
}
