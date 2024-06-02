// Package masterLoginBonusSchedule ログインボーナススケジュール
package masterLoginBonusSchedule

func SetMasterLoginBonusSchedule(masterLoginBonusScheduleId int64, masterLoginBonusId int64, step int32, name string) *MasterLoginBonusSchedule {
	return &MasterLoginBonusSchedule{
		MasterLoginBonusScheduleId: masterLoginBonusScheduleId,
		MasterLoginBonusId:         masterLoginBonusId,
		Step:                       step,
		Name:                       name,
	}
}
