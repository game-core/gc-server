// Package masterActionStep アクションステップ
package masterActionStep

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionStep"
)

type MasterActionSteps []*MasterActionStep

type MasterActionStep struct {
	MasterActionStepId   int64
	Name                 string
	MasterActionStepEnum masterActionStep.MasterActionStepEnum
}

func NewMasterActionStep() *MasterActionStep {
	return &MasterActionStep{}
}

func NewMasterActionSteps() MasterActionSteps {
	return MasterActionSteps{}
}

func SetMasterActionStep(masterActionStepId int64, name string, masterActionStepEnum masterActionStep.MasterActionStepEnum) *MasterActionStep {
	return &MasterActionStep{
		MasterActionStepId:   masterActionStepId,
		Name:                 name,
		MasterActionStepEnum: masterActionStepEnum,
	}
}

func (t *MasterActionStep) TableName() string {
	return "master_action_step"
}
