//go:generate mockgen -source=./transaction_service.go -destination=./transaction_service_mock.gen.go -package=transaction
package transaction

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/adminTransaction"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/commonTransaction"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/masterTransaction"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/userTransaction"
)

type TransactionService interface {
	AdminMysqlBegin(ctx context.Context) (*gorm.DB, error)
	AdminMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	CommonMysqlBegin(ctx context.Context) (*gorm.DB, error)
	CommonMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	MasterMysqlBegin(ctx context.Context) (*gorm.DB, error)
	MasterMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	UserMysqlBegin(ctx context.Context, userId string) (*gorm.DB, error)
	UserMysqlEnd(ctx context.Context, tx *gorm.DB, err error)
	MultiUserMysqlBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error)
	MultiUserMysqlEnd(ctx context.Context, txs map[string]*gorm.DB, err error)
	UserRedisBegin() redis.Pipeliner
	UserRedisEnd(ctx context.Context, tx redis.Pipeliner, err error)
}

type transactionService struct {
	adminTransactionMysqlRepository  adminTransaction.AdminTransactionMysqlRepository
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository
	userTransactionMysqlRepository   userTransaction.UserTransactionMysqlRepository
	userTransactionRedisRepository   userTransaction.UserTransactionRedisRepository
}

func NewTransactionService(
	adminTransactionMysqlRepository adminTransaction.AdminTransactionMysqlRepository,
	commonTransactionMysqlRepository commonTransaction.CommonTransactionMysqlRepository,
	masterTransactionMysqlRepository masterTransaction.MasterTransactionMysqlRepository,
	userTransactionMysqlRepository userTransaction.UserTransactionMysqlRepository,
	userTransactionRedisRepository userTransaction.UserTransactionRedisRepository,
) TransactionService {
	return &transactionService{
		adminTransactionMysqlRepository:  adminTransactionMysqlRepository,
		commonTransactionMysqlRepository: commonTransactionMysqlRepository,
		masterTransactionMysqlRepository: masterTransactionMysqlRepository,
		userTransactionMysqlRepository:   userTransactionMysqlRepository,
		userTransactionRedisRepository:   userTransactionRedisRepository,
	}
}

// AdminMysqlBegin トランザクションを開始する
func (s *transactionService) AdminMysqlBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.adminTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// AdminMysqlEnd トランザクションを終了する
func (s *transactionService) AdminMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.adminTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.adminTransactionMysqlRepository.Rollback", err)
		}
	} else {
		if err := s.adminTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.adminTransactionMysqlRepository.Commit", err)
		}
	}
}

// CommonMysqlBegin トランザクションを開始する
func (s *transactionService) CommonMysqlBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.commonTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

// CommonMysqlEnd トランザクションを終了する
func (s *transactionService) CommonMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.commonTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.commonTransactionMysqlRepository.Rollback", err)
		}
	} else {
		if err := s.commonTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.commonTransactionMysqlRepository.Commit", err)
		}
	}
}

// MasterMysqlBegin トランザクションを開始する
func (s *transactionService) MasterMysqlBegin(ctx context.Context) (*gorm.DB, error) {
	tx, err := s.masterTransactionMysqlRepository.Begin(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.masterTransactionMysqlRepository.Begin", err)
	}

	return tx, nil
}

// MasterMysqlEnd トランザクションを終了する
func (s *transactionService) MasterMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.masterTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.masterTransactionMysqlRepository.Rollback", err)
		}
	} else {
		if err := s.masterTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.masterTransactionMysqlRepository.Commit", err)
		}
	}
}

// UserMysqlBegin トランザクションを開始する
func (s *transactionService) UserMysqlBegin(ctx context.Context, userId string) (*gorm.DB, error) {
	tx, err := s.userTransactionMysqlRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
	if err != nil {
		return nil, errors.NewMethodError("s.userTransactionMysqlRepository.Begin", err)
	}

	return tx, nil
}

// UserMysqlEnd トランザクションを終了する
func (s *transactionService) UserMysqlEnd(ctx context.Context, tx *gorm.DB, err error) {
	if err != nil {
		if err := s.userTransactionMysqlRepository.Rollback(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.userTransactionMysqlRepository.Rollback", err)
		}
	} else {
		if err := s.userTransactionMysqlRepository.Commit(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.userTransactionMysqlRepository.Commit", err)
		}
	}
}

// MultiUserMysqlBegin マルチトランザクションを開始する
func (s *transactionService) MultiUserMysqlBegin(ctx context.Context, userIds []string) (map[string]*gorm.DB, error) {
	txs := make(map[string]*gorm.DB)
	for _, userId := range userIds {
		tx, err := s.userTransactionMysqlRepository.Begin(ctx, keys.GetShardKeyByUserId(userId))
		if err != nil {
			return nil, errors.NewMethodError("s.userTransactionMysqlRepository.Begin", err)
		}

		txs[userId] = tx
	}

	return txs, nil
}

// MultiUserMysqlEnd マルチトランザクションを終了する
func (s *transactionService) MultiUserMysqlEnd(ctx context.Context, txs map[string]*gorm.DB, err error) {
	if err != nil {
		for _, tx := range txs {
			if rollbackErr := s.userTransactionMysqlRepository.Rollback(ctx, tx); rollbackErr != nil {
				errors.NewMethodErrorLog("s.userTransactionMysqlRepository.Rollback", rollbackErr)
			}
		}
		return
	}

	for _, tx := range txs {
		if commitErr := s.userTransactionMysqlRepository.Commit(ctx, tx); commitErr != nil {
			errors.NewMethodErrorLog("s.userTransactionMysqlRepository.Commit", commitErr)
		}
	}
}

// UserRedisBegin トランザクションを開始する
func (s *transactionService) UserRedisBegin() redis.Pipeliner {
	return s.userTransactionRedisRepository.Begin()
}

// UserRedisEnd トランザクションを終了する
func (s *transactionService) UserRedisEnd(ctx context.Context, tx redis.Pipeliner, err error) {
	if err != nil {
		s.userTransactionRedisRepository.Discard(tx)
		errors.NewMethodErrorLog("s.userTransactionRedisRepository.Discard", err)
	} else {
		if err := s.userTransactionRedisRepository.Commit(ctx, tx); err != nil {
			errors.NewMethodErrorLog("s.userTransactionRedisRepository.Commit", err)
		}
	}
}
