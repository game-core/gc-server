// Package masterResource リソース
package masterResource

type MasterResources []*MasterResource

type MasterResource struct {
	MasterResourceId   int64
	Name               string
	MasterResourceEnum MasterResourceEnum
}

func NewMasterResource() *MasterResource {
	return &MasterResource{}
}

func NewMasterResources() MasterResources {
	return MasterResources{}
}

func SetMasterResource(masterResourceId int64, name string, masterResourceEnum MasterResourceEnum) *MasterResource {
	return &MasterResource{
		MasterResourceId:   masterResourceId,
		Name:               name,
		MasterResourceEnum: masterResourceEnum,
	}
}
