// Package masterRarity レアリティ
package masterRarity

type MasterRarities []*MasterRarity

type MasterRarity struct {
	MasterRarityId   int64
	Name             string
	MasterRarityEnum MasterRarityEnum
}

func NewMasterRarity() *MasterRarity {
	return &MasterRarity{}
}

func NewMasterRarities() MasterRarities {
	return MasterRarities{}
}

func SetMasterRarity(masterRarityId int64, name string, masterRarityEnum MasterRarityEnum) *MasterRarity {
	return &MasterRarity{
		MasterRarityId:   masterRarityId,
		Name:             name,
		MasterRarityEnum: masterRarityEnum,
	}
}
