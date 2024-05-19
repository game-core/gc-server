// Package masterResource リソース
package masterResource

import (
	"github.com/game-core/gc-server/pkg/domain/model/resource/masterResource"
)

type MasterResources []*MasterResource

type MasterResource struct {
	MasterResourceId   int64
	Name               string
	MasterResourceEnum masterResource.MasterResourceEnum
}

func NewMasterResource() *MasterResource {
	return &MasterResource{}
}

func NewMasterResources() MasterResources {
	return MasterResources{}
}

func SetMasterResource(masterResourceId int64, name string, masterResourceEnum masterResource.MasterResourceEnum) *MasterResource {
	return &MasterResource{
		MasterResourceId:   masterResourceId,
		Name:               name,
		MasterResourceEnum: masterResourceEnum,
	}
}

func (t *MasterResource) TableName() string {
	return "master_resource"
}
