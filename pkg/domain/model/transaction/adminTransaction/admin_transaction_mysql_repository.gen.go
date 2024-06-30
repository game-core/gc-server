//go:generate mockgen -source=./admin_transaction_mysql_repository.gen.go -destination=./admin_transaction_mysql_repository_mock.gen.go -package=adminTransaction
package adminTransaction

import (
	"context"

	"gorm.io/gorm"
)

type AdminTransactionMysqlRepository interface {
	Begin(ctx context.Context) (*gorm.DB, error)
	Commit(ctx context.Context, tx *gorm.DB) error
	Rollback(ctx context.Context, tx *gorm.DB) error
}
