// Package masterActionStep アクションステップ
package masterActionStep

type MasterActionSteps []*MasterActionStep

type MasterActionStep struct {
	MasterActionStepId   int64
	Name                 string
	MasterActionStepEnum MasterActionStepEnum
}

func NewMasterActionStep() *MasterActionStep {
	return &MasterActionStep{}
}

func NewMasterActionSteps() MasterActionSteps {
	return MasterActionSteps{}
}

func SetMasterActionStep(masterActionStepId int64, name string, masterActionStepEnum MasterActionStepEnum) *MasterActionStep {
	return &MasterActionStep{
		MasterActionStepId:   masterActionStepId,
		Name:                 name,
		MasterActionStepEnum: masterActionStepEnum,
	}
}
