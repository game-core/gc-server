// Package masterAction アクション
package masterAction

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
)

type MasterActions []*MasterAction

type MasterAction struct {
	MasterActionId          int64
	Name                    string
	MasterActionStepEnum    masterActionStep.MasterActionStepEnum
	MasterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum
	TargetId                *int64
	TriggerMasterActionId   *int64
	Expiration              *int32
}

func NewMasterAction() *MasterAction {
	return &MasterAction{}
}

func NewMasterActions() MasterActions {
	return MasterActions{}
}

func SetMasterAction(masterActionId int64, name string, masterActionStepEnum masterActionStep.MasterActionStepEnum, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum, targetId *int64, triggerMasterActionId *int64, expiration *int32) *MasterAction {
	return &MasterAction{
		MasterActionId:          masterActionId,
		Name:                    name,
		MasterActionStepEnum:    masterActionStepEnum,
		MasterActionTriggerEnum: masterActionTriggerEnum,
		TargetId:                targetId,
		TriggerMasterActionId:   triggerMasterActionId,
		Expiration:              expiration,
	}
}
