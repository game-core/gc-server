// Package masterActionTrigger アクショントリガー
package masterActionTrigger

type MasterActionTriggers []*MasterActionTrigger

type MasterActionTrigger struct {
	MasterActionTriggerId   int64
	Name                    string
	MasterActionTriggerEnum MasterActionTriggerEnum
}

func NewMasterActionTrigger() *MasterActionTrigger {
	return &MasterActionTrigger{}
}

func NewMasterActionTriggers() MasterActionTriggers {
	return MasterActionTriggers{}
}

func SetMasterActionTrigger(masterActionTriggerId int64, name string, masterActionTriggerEnum MasterActionTriggerEnum) *MasterActionTrigger {
	return &MasterActionTrigger{
		MasterActionTriggerId:   masterActionTriggerId,
		Name:                    name,
		MasterActionTriggerEnum: masterActionTriggerEnum,
	}
}
