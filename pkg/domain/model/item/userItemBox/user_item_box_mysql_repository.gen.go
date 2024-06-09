// Package userItemBox ユーザーアイテムボックス
//
//go:generate mockgen -source=./user_item_box_mysql_repository.gen.go -destination=./user_item_box_mysql_repository_mock.gen.go -package=userItemBox
package userItemBox

import (
	"context"
	"time"

	"github.com/game-core/gc-server/config/logger"
)

type UserItemBoxCloudWatchRepository interface {
	Create(ctx context.Context, now time.Time, level logger.LogLevel, m *UserItemBox)
	CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms UserItemBoxes)
}
