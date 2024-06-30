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
	"github.com/game-core/gc-server/internal/pointers"
	"github.com/game-core/gc-server/pkg/domain/model/item/userItemBox"
)

type userItemBoxCloudWatchDao struct {
	ReadCloudWatchConn  *cloudwatchlogs.Client
	WriteCloudWatchConn *cloudwatchlogs.Client
}

func NewUserItemBoxCloudWatchDao(conn *logger.CloudWatchHandler) userItemBox.UserItemBoxCloudWatchRepository {
	return &userItemBoxCloudWatchDao{
		ReadCloudWatchConn:  conn.User.ReadCloudWatchConn,
		WriteCloudWatchConn: conn.User.WriteCloudWatchConn,
	}
}

func (s *userItemBoxCloudWatchDao) Create(ctx context.Context, now time.Time, level logger.LogLevel, m *userItemBox.UserItemBox) {
	timestamp := now.Unix() * 1000
	t := &UserItemBox{
		UserId:       m.UserId,
		MasterItemId: m.MasterItemId,
		Count:        m.Count,
	}
	message := string(logger.SetLogMessage(now, level, t).ToJson())

	switch os.Getenv("APP_ENV") {
	case "prod":
		if err := s.creteToCloudWatch(ctx, timestamp, message); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	case "dev":
		if err := s.creteToFile("./log/gc_server_user.log", fmt.Sprintf("%s %s\n", now.Format(time.RFC3339), message)); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	default:
		errors.NewErrorLog("APP_ENV is invalid")
	}
}

func (s *userItemBoxCloudWatchDao) CreateList(ctx context.Context, now time.Time, level logger.LogLevel, ms userItemBox.UserItemBoxes) {
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

	switch os.Getenv("APP_ENV") {
	case "prod":
		if err := s.creteToCloudWatch(ctx, timestamp, message); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	case "dev":
		if err := s.creteToFile("./log/gc_server_user.log", fmt.Sprintf("%s %s\n", now.Format(time.RFC3339), message)); err != nil {
			errors.NewMethodErrorLog("appendToFile", err)
		}
	default:
		errors.NewErrorLog("APP_ENV is invalid")
	}
}

func (s *userItemBoxCloudWatchDao) creteToCloudWatch(ctx context.Context, timestamp int64, message string) error {
	if _, err := s.WriteCloudWatchConn.PutLogEvents(
		ctx,
		&cloudwatchlogs.PutLogEventsInput{
			LogEvents: []types.InputLogEvent{
				{
					Timestamp: &timestamp,
					Message:   &message,
				},
			},
			LogGroupName:  pointers.StringToPointer(os.Getenv("USER_LOG_GROUP_NAME")),
			LogStreamName: pointers.StringToPointer(os.Getenv("USER_LOG_STREAM_NAME")),
		},
	); err != nil {
		return errors.NewMethodError("s.WriteCloudWatchConn.PutLogEvents", err)
	}

	return nil
}

func (s *userItemBoxCloudWatchDao) creteToFile(fileName, message string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.NewMethodError("os.OpenFile", err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			errors.NewMethodErrorLog("f.Close", err)
		}
	}(f)
	if _, err := f.WriteString(message); err != nil {
		return errors.NewMethodError("f.WriteString", err)
	}

	return nil
}
