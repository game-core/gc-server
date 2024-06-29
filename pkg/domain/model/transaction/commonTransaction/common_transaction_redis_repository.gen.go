//go:generate mockgen -source=./common_transaction_redis_repository.gen.go -destination=./common_transaction_redis_repository_mock.gen.go -package=commonTransaction
package commonTransaction

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CommonTransactionRedisRepository interface {
	Begin() redis.Pipeliner
	Commit(ctx context.Context, tx redis.Pipeliner) error
	Discard(tx redis.Pipeliner)
}
