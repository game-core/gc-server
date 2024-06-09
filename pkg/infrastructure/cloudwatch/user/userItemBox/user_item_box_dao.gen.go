// Package userItemBox ユーザーアイテムボックス
package userItemBox

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"

	"github.com/game-core/gc-server/config/logger"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type userItemBoxDao struct {
	ReadCloudWatchConn  *cloudwatchlogs.Client
	WriteCloudWatchConn *cloudwatchlogs.Client
}

func NewUserItemBoxDao(conn *logger.CloudWatchHandler) userItemBox.UserItemBoxCloudWatchRepository {
	return &userItemBoxDao{
		ReadCloudWatchConn:  conn.User.ReadCloudWatchConn,
		WriteCloudWatchConn: conn.User.WriteCloudWatchConn,
	}
}

func (s *userItemBoxDao) Create(ctx context.Context, now time.Time, level logger.LogLevel, m *userItemBox.UserItemBox) {
	timestamp := now.Unix() * 1000
	t := &UserItemBox{
		UserId:       m.UserId,
		MasterItemId: m.MasterItemId,
		Count:        m.Count,
	}
	message := string(logger.SetLogMessage(now, level, t).ToJson())

	if os.Getenv("APP_ENV") == "prod" {
		if err := s.creteToCloudWatch(ctx, timestamp, os.Getenv("USER_LOG_GROUP_NAME"), os.Getenv("USER_LOG_STREAM_NAME"), message); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	} else if os.Getenv("APP_ENV") == "dev" {
		if err := s.creteToFile("./log/gc_server_user.log", fmt.Sprintf("%s %s\n", now.Format(time.RFC3339), message)); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	}
}

func (s *userItemBoxDao) CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms userItemBox.UserItemBoxes) {
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

	if os.Getenv("APP_ENV") == "prod" {
		if err := s.creteToCloudWatch(ctx, timestamp, os.Getenv("USER_LOG_GROUP_NAME"), os.Getenv("USER_LOG_STREAM_NAME"), message); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	} else if os.Getenv("APP_ENV") == "dev" {
		if err := s.creteToFile("./log/gc_server_user.log", fmt.Sprintf("%s %s\n", now.Format(time.RFC3339), message)); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	}
}

func (s *userItemBoxDao) creteToCloudWatch(ctx context.Context, timestamp int64, logGroupName, logStreamName, message string) error {
	if _, err := s.WriteCloudWatchConn.PutLogEvents(
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
		errors.NewMethodErrorLog("s.WriteCloudWatchConn.PutLogEvents", err)
	}

	return nil
}

func (s *userItemBoxDao) creteToFile(fileName, message string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			errors.NewMethodErrorLog("f.Close", err)
		}
	}(f)
	if _, err := f.WriteString(message); err != nil {
		return err
	}

	return nil
}
