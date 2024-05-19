// Package masterRarity レアリティ
package masterRarity

import (
	"github.com/game-core/gc-server/pkg/domain/model/rarity/masterRarity"
)

type MasterRarities []*MasterRarity

type MasterRarity struct {
	MasterRarityId   int64
	Name             string
	MasterRarityEnum masterRarity.MasterRarityEnum
}

func NewMasterRarity() *MasterRarity {
	return &MasterRarity{}
}

func NewMasterRarities() MasterRarities {
	return MasterRarities{}
}

func SetMasterRarity(masterRarityId int64, name string, masterRarityEnum masterRarity.MasterRarityEnum) *MasterRarity {
	return &MasterRarity{
		MasterRarityId:   masterRarityId,
		Name:             name,
		MasterRarityEnum: masterRarityEnum,
	}
}

func (t *MasterRarity) TableName() string {
	return "master_rarity"
}
