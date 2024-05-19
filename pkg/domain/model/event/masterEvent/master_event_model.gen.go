// Package masterEvent イベント
package masterEvent

import (
	"time"
)

type MasterEvents []*MasterEvent

type MasterEvent struct {
	MasterEventId int64
	Name          string
	ResetHour     int32
	IntervalHour  int32
	RepeatSetting bool
	StartAt       time.Time
	EndAt         *time.Time
}

func NewMasterEvent() *MasterEvent {
	return &MasterEvent{}
}

func NewMasterEvents() MasterEvents {
	return MasterEvents{}
}

func SetMasterEvent(masterEventId int64, name string, resetHour int32, intervalHour int32, repeatSetting bool, startAt time.Time, endAt *time.Time) *MasterEvent {
	return &MasterEvent{
		MasterEventId: masterEventId,
		Name:          name,
		ResetHour:     resetHour,
		IntervalHour:  intervalHour,
		RepeatSetting: repeatSetting,
		StartAt:       startAt,
		EndAt:         endAt,
	}
}
