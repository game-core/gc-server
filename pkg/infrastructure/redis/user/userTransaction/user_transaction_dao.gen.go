package masterTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/userTransaction"
)

type userTransactionRedisDao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func NewUserTransactionRedisDao(conn *database.RedisHandler) userTransaction.UserTransactionRedisRepository {
	return &userTransactionRedisDao{
		ReadRedisConn:  conn.User.ReadRedisConn,
		WriteRedisConn: conn.User.WriteRedisConn,
	}
}

func (d *userTransactionRedisDao) Begin() redis.Pipeliner {
	return d.WriteRedisConn.TxPipeline()
}

func (d *userTransactionRedisDao) Commit(ctx context.Context, tx redis.Pipeliner) error {
	if _, err := tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (d *userTransactionRedisDao) Rollback(tx redis.Pipeliner) {
	tx.Discard()
}
