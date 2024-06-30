package adminTransaction

import (
	"context"

	"gorm.io/gorm"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/pkg/domain/model/transaction/adminTransaction"
)

type adminTransactionMysqlDao struct {
	ReadMysqlConn  *gorm.DB
	WriteMysqlConn *gorm.DB
}

func NewAdminTransactionMysqlDao(conn *database.MysqlHandler) adminTransaction.AdminTransactionMysqlRepository {
	return &adminTransactionMysqlDao{
		ReadMysqlConn:  conn.Admin.ReadMysqlConn,
		WriteMysqlConn: conn.Admin.WriteMysqlConn,
	}
}

func (d *adminTransactionMysqlDao) Begin(ctx context.Context) (*gorm.DB, error) {
	tx := d.WriteMysqlConn.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, err
	}

	return tx, nil
}

func (d *adminTransactionMysqlDao) Commit(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Commit()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (d *adminTransactionMysqlDao) Rollback(ctx context.Context, tx *gorm.DB) error {
	tx.WithContext(ctx).Rollback()
	if err := tx.Error; err != nil {
		return err
	}

	return nil
}
