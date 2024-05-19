// Package action アクション実行リクエスト
package action

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
)

type ActionRunRequests []*ActionRunRequest

type ActionRunRequest struct {
	UserId               string
	MasterActionStepEnum masterActionStep.MasterActionStepEnum
	TargetId             *int64
}

func NewActionRunRequest() *ActionRunRequest {
	return &ActionRunRequest{}
}

func NewActionRunRequests() ActionRunRequests {
	return ActionRunRequests{}
}

func SetActionRunRequest(userId string, masterActionStepEnum masterActionStep.MasterActionStepEnum, targetId *int64) *ActionRunRequest {
	return &ActionRunRequest{
		UserId:               userId,
		MasterActionStepEnum: masterActionStepEnum,
		TargetId:             targetId,
	}
}
