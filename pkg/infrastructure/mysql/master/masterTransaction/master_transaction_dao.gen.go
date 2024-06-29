package masterTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/masterTransaction"
)

type masterTransactionMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewMasterTransactionMysqlDao(conn *database.MysqlHandler) masterTransaction.MasterTransactionMysqlRepository {
	return &masterTransactionMysqlDao{
		ReadMysqlConn:  conn.Master.ReadMysqlConn,
		WriteMysqlConn: conn.Master.WriteMysqlConn,
	}
}

func (d *masterTransactionMysqlDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteMysqlConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *masterTransactionMysqlDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *masterTransactionMysqlDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
