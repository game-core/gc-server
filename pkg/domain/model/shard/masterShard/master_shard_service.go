package masterShard

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
)

var MasterShardInstances MasterShards

// GetShardKey シャードキーを取得する
func (s *MasterShards) GetShardKey(ctx context.Context, masterShardMysqlRepository MasterShardMysqlRepository) (string, error) {
	if len(MasterShardInstances) <= 0 {
		masterShards, err := masterShardMysqlRepository.FindList(ctx)
		if err != nil {
			return "", errors.NewMethodError("s.masterShardMysqlRepository.FindList", err)
		}
		if len(masterShards) <= 0 {
			return "", errors.NewError("common_shard does not exist")
		}

		MasterShardInstances = masterShards
	}

	minShard := MasterShardInstances[0]
	minIndex := 0
	for i, s := range MasterShardInstances {
		if s.Count < minShard.Count {
			minShard = s
			minIndex = i
			break
		}
	}

	minShard.Count++
	MasterShardInstances[minIndex] = minShard

	return minShard.ShardKey, nil
}
