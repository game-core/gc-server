// Package masterItem アイテム
package masterItem

import (
	"github.com/game-core/gc-server/pkg/domain/model/rarity/masterRarity"
	"github.com/game-core/gc-server/pkg/domain/model/resource/masterResource"
)

type MasterItems []*MasterItem

type MasterItem struct {
	MasterItemId       int64
	Name               string
	MasterResourceEnum masterResource.MasterResourceEnum
	MasterRarityEnum   masterRarity.MasterRarityEnum
	Content            string
}

func NewMasterItem() *MasterItem {
	return &MasterItem{}
}

func NewMasterItems() MasterItems {
	return MasterItems{}
}

func SetMasterItem(masterItemId int64, name string, masterResourceEnum masterResource.MasterResourceEnum, masterRarityEnum masterRarity.MasterRarityEnum, content string) *MasterItem {
	return &MasterItem{
		MasterItemId:       masterItemId,
		Name:               name,
		MasterResourceEnum: masterResourceEnum,
		MasterRarityEnum:   masterRarityEnum,
		Content:            content,
	}
}
