// Package masterShard シャード
//
//go:generate mockgen -source=./master_shard_mysql_repository.gen.go -destination=./master_shard_mysql_repository_mock.gen.go -package=masterShard
package masterShard

import (
	"context"

	"gorm.io/gorm"
)

type MasterShardMysqlRepository interface {
	Find(ctx context.Context, masterShardId int64) (*MasterShard, error)
	FindOrNil(ctx context.Context, masterShardId int64) (*MasterShard, error)
	FindByShardKey(ctx context.Context, shardKey string) (*MasterShard, error)
	FindOrNilByShardKey(ctx context.Context, shardKey string) (*MasterShard, error)
	FindList(ctx context.Context) (MasterShards, error)
	FindListByShardKey(ctx context.Context, shardKey string) (MasterShards, error)
	Create(ctx context.Context, tx *gorm.DB, m *MasterShard) (*MasterShard, error)
	CreateList(ctx context.Context, tx *gorm.DB, ms MasterShards) (MasterShards, error)
	Update(ctx context.Context, tx *gorm.DB, m *MasterShard) (*MasterShard, error)
	Delete(ctx context.Context, tx *gorm.DB, m *MasterShard) error
}
