// Package action アクション確認リクエスト
package action

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
)

type ActionCheckRequests []*ActionCheckRequest

type ActionCheckRequest struct {
	UserId               string
	MasterActionStepEnum masterActionStep.MasterActionStepEnum
	TargetId             *int64
}

func NewActionCheckRequest() *ActionCheckRequest {
	return &ActionCheckRequest{}
}

func NewActionCheckRequests() ActionCheckRequests {
	return ActionCheckRequests{}
}

func SetActionCheckRequest(userId string, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) *ActionCheckRequest {
	return &ActionCheckRequest{
		UserId:               userId,
		MasterActionStepEnum: masterActionStepEnum,
		TargetId:             targetId,
	}
}
