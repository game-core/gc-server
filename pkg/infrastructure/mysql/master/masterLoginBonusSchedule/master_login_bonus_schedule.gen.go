// Package masterLoginBonusSchedule ログインボーナススケジュール
package masterLoginBonusSchedule

type MasterLoginBonusSchedules []*MasterLoginBonusSchedule

type MasterLoginBonusSchedule struct {
	MasterLoginBonusScheduleId int64
	MasterLoginBonusId         int64
	Step                       int32
	Name                       string
}

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

func (t *MasterLoginBonusSchedule) TableName() string {
	return "master_login_bonus_schedule"
}
