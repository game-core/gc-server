package userTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/userTransaction"
)

type userTransactionMysqlDao struct {
	ShardMysqlConn *database.ShardMysqlConn
}

func NewUserTransactionMysqlDao(conn *database.MysqlHandler) userTransaction.UserTransactionMysqlRepository {
	return &userTransactionMysqlDao{
		ShardMysqlConn: conn.User,
	}
}

func (d *userTransactionMysqlDao) Begin(ctx context.Context, shardKey string) (*gorm.DB, error) {
	tx := d.ShardMysqlConn.Shards[shardKey].WriteMysqlConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *userTransactionMysqlDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *userTransactionMysqlDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
