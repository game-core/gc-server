// Package masterShard シャード
package masterShard

type MasterShards []*MasterShard

type MasterShard struct {
	MasterShardId int64
	ShardKey      string
	Name          string
}

func NewMasterShard() *MasterShard {
	return &MasterShard{}
}

func NewMasterShards() MasterShards {
	return MasterShards{}
}

func SetMasterShard(masterShardId int64, shardKey string, name string) *MasterShard {
	return &MasterShard{
		MasterShardId: masterShardId,
		ShardKey:      shardKey,
		Name:          name,
	}
}

func (t *MasterShard) TableName() string {
	return "master_shard"
}
