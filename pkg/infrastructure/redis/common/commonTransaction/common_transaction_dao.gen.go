package masterTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/commonTransaction"
)

type commonTransactionDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewCommonTransactionDao(conn *database.RedisHandler) commonTransaction.CommonTransactionRedisRepository {
	return &commonTransactionDao{
		ReadRedisConn:  conn.Common.ReadRedisConn,
		WriteRedisConn: conn.Common.WriteRedisConn,
	}
}

func (d *commonTransactionDao) Begin() redis.Pipeliner {
	return d.WriteRedisConn.TxPipeline()
}

func (d *commonTransactionDao) Commit(ctx context.Context, tx redis.Pipeliner) error {
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (d *commonTransactionDao) Rollback(tx redis.Pipeliner) {
	tx.Discard()
}
