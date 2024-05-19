// Package action アクションマスター取得レスポンス
package action

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterAction"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionRun"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
)

type ActionGetListMasterResponses []*ActionGetListMasterResponse

type ActionGetListMasterResponse struct {
	MasterActions        masterAction.MasterActions
	MasterActionRuns     masterActionRun.MasterActionRuns
	MasterActionSteps    masterActionStep.MasterActionSteps
	MasterActionTriggers masterActionTrigger.MasterActionTriggers
}

func NewActionGetListMasterResponse() *ActionGetListMasterResponse {
	return &ActionGetListMasterResponse{}
}

func NewActionGetListMasterResponses() ActionGetListMasterResponses {
	return ActionGetListMasterResponses{}
}

func SetActionGetListMasterResponse(masterActions masterAction.MasterActions, masterActionRuns masterActionRun.MasterActionRuns, masterActionSteps masterActionStep.MasterActionSteps, masterActionTriggers masterActionTrigger.MasterActionTriggers) *ActionGetListMasterResponse {
	return &ActionGetListMasterResponse{
		MasterActions:        masterActions,
		MasterActionRuns:     masterActionRuns,
		MasterActionSteps:    masterActionSteps,
		MasterActionTriggers: masterActionTriggers,
	}
}
