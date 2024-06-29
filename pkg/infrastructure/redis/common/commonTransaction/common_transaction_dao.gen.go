package masterTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/commonTransaction"
)

type commonTransactionRedisDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewCommonTransactionRedisDao(conn *database.RedisHandler) commonTransaction.CommonTransactionRedisRepository {
	return &commonTransactionRedisDao{
		ReadRedisConn:  conn.Common.ReadRedisConn,
		WriteRedisConn: conn.Common.WriteRedisConn,
	}
}

func (d *commonTransactionRedisDao) Begin() redis.Pipeliner {
	return d.WriteRedisConn.TxPipeline()
}

func (d *commonTransactionRedisDao) Commit(ctx context.Context, tx redis.Pipeliner) error {
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (d *commonTransactionRedisDao) Rollback(tx redis.Pipeliner) {
	tx.Discard()
}
