// Package masterLoginBonusSchedule ログインボーナススケジュール
package masterLoginBonusSchedule

type MasterLoginBonusSchedules []*MasterLoginBonusSchedule

func NewMasterLoginBonusSchedule() *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{}
}

func NewMasterLoginBonusSchedules() MasterLoginBonusSchedules {
	return MasterLoginBonusSchedules{}
}

func SetMasterLoginBonusSchedule(masterLoginBonusScheduleId int64, masterLoginBonusId int64, step int32, name string) *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterLoginBonusId:         masterLoginBonusId,
		Step:                       step,
		Name:                       name,
	}
}
