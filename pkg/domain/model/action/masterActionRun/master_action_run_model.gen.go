// Package masterActionRun 実行されるアクション
package masterActionRun

type MasterActionRuns []*MasterActionRun

type MasterActionRun struct {
	MasterActionRunId int64
	Name              string
	MasterActionId    int64
}

func NewMasterActionRun() *MasterActionRun {
	return &MasterActionRun{}
}

func NewMasterActionRuns() MasterActionRuns {
	return MasterActionRuns{}
}

func SetMasterActionRun(masterActionRunId int64, name string, masterActionId int64) *MasterActionRun {
	return &MasterActionRun{
		MasterActionRunId: masterActionRunId,
		Name:              name,
		MasterActionId:    masterActionId,
	}
}
