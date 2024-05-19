//go:generate mockgen -source=./shard_service.go -destination=./shard_service_mock.gen.go -package=shard
package shard

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/shard/masterShard"
)

type ShardService interface {
	GetShardKey(ctx context.Context) (string, error)
}

type shardService struct {
	masterShardMysqlRepository masterShard.MasterShardMysqlRepository
}

func NewShardService(
	masterShardMysqlRepository masterShard.MasterShardMysqlRepository,
) ShardService {
	return &shardService{
		masterShardMysqlRepository: masterShardMysqlRepository,
	}
}

// GetShardKey シャードキーを取得して更新する
func (s *shardService) GetShardKey(ctx context.Context) (string, error) {
	masterShards := masterShard.NewMasterShards()
	shardKey, err := masterShards.GetShardKey(ctx, s.masterShardMysqlRepository)
	if err != nil {
		return "", errors.NewMethodError("shards.GetShardKey", err)
	}

	return shardKey, nil
}
