// Package userItemBox ユーザーアイテムボックス
package userItemBox

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/game-core/gc-server/config/logger"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type userItemBoxDao struct {
	ReadMysqlConn  *cloudwatchlogs.Client
	WriteMysqlConn *cloudwatchlogs.Client
}

func NewUserItemBoxDao(conn *logger.CloudWatchHandler) userItemBox.UserItemBoxCloudWatchRepository {
	return &userItemBoxDao{
		ReadMysqlConn:  conn.User.ReadMysqlConn,
		WriteMysqlConn: conn.User.WriteMysqlConn,
	}
}

func (s *userItemBoxDao) Create(ctx context.Context, now time.Time, level logger.LogLevel, m *userItemBox.UserItemBox) {
	logGroupName := os.Getenv("USER_LOG_GROUP_NAME")
	logStreamName := os.Getenv("USER_LOG_STREAM_NAME")
	timestamp := now.Unix() * 1000
	t := &UserItemBox{
		UserId:       m.UserId,
		MasterItemId: m.MasterItemId,
		Count:        m.Count,
	}
	message := string(logger.SetLogMessage(now, level, t).ToJson())

	if _, err := s.WriteMysqlConn.PutLogEvents(
		ctx,
		&cloudwatchlogs.PutLogEventsInput{
			LogEvents: []types.InputLogEvent{
				{
					Timestamp: &timestamp,
					Message:   &message,
				},
			},
			LogGroupName:  &logGroupName,
			LogStreamName: &logStreamName,
		},
	); err != nil {
		errors.NewMethodErrorLog("s.WriteMysqlConn.PutLogEvents", err)
	}
}

func (s *userItemBoxDao) CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms userItemBox.UserItemBoxes) {
	logGroupName := os.Getenv("USER_LOG_GROUP_NAME")
	logStreamName := os.Getenv("USER_LOG_STREAM_NAME")
	timestamp := now.Unix() * 1000
	ts := NewUserItemBoxes()
	for _, m := range ms {
		t := &UserItemBox{
			UserId:       m.UserId,
			MasterItemId: m.MasterItemId,
			Count:        m.Count,
		}
		ts = append(ts, t)
	}
	message := string(logger.SetLogMessage(now, level, ts).ToJson())

	if _, err := s.WriteMysqlConn.PutLogEvents(
		ctx,
		&cloudwatchlogs.PutLogEventsInput{
			LogEvents: []types.InputLogEvent{
				{
					Timestamp: &timestamp,
					Message:   &message,
				},
			},
			LogGroupName:  &logGroupName,
			LogStreamName: &logStreamName,
		},
	); err != nil {
		errors.NewMethodErrorLog("s.WriteMysqlConn.PutLogEvents", err)
	}
}
