// Package masterShard シャード
package masterShard

type MasterShards []*MasterShard

type MasterShard struct {
	MasterShardId int64
	Name          string
	ShardKey      string
	Count         int32
}

func NewMasterShard() *MasterShard {
	return &MasterShard{}
}

func NewMasterShards() MasterShards {
	return MasterShards{}
}

func SetMasterShard(masterShardId int64, name string, shardKey string, count int32) *MasterShard {
	return &MasterShard{
		MasterShardId: masterShardId,
		Name:          name,
		ShardKey:      shardKey,
		Count:         count,
	}
}

func (t *MasterShard) TableName() string {
	return "master_shard"
}
