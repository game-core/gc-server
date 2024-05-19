// Package masterActionTrigger アクショントリガー
package masterActionTrigger

import (
	"github.com/game-core/gc-server/pkg/domain/model/action/masterActionTrigger"
)

type MasterActionTriggers []*MasterActionTrigger

type MasterActionTrigger struct {
	MasterActionTriggerId   int64
	Name                    string
	MasterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum
}

func NewMasterActionTrigger() *MasterActionTrigger {
	return &MasterActionTrigger{}
}

func NewMasterActionTriggers() MasterActionTriggers {
	return MasterActionTriggers{}
}

func SetMasterActionTrigger(masterActionTriggerId int64, name string, masterActionTriggerEnum masterActionTrigger.MasterActionTriggerEnum) *MasterActionTrigger {
	return &MasterActionTrigger{
		MasterActionTriggerId:   masterActionTriggerId,
		Name:                    name,
		MasterActionTriggerEnum: masterActionTriggerEnum,
	}
}

func (t *MasterActionTrigger) TableName() string {
	return "master_action_trigger"
}
