package logger

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

	"github.com/game-core/gc-server/internal/errors"
)

var CloudWatchHandlerInstance *CloudWatchHandler

type CloudWatchHandler struct {
	User *CloudWatchConn
}

type CloudWatchConn struct {
	ReadMysqlConn  *cloudwatchlogs.Client
	WriteMysqlConn *cloudwatchlogs.Client
}

// NewCloudWatch インスタンスを作成する
func NewCloudWatch() *CloudWatchHandler {
	return CloudWatchHandlerInstance
}

// InitCloudWatch 初期化する
func InitCloudWatch() (*CloudWatchHandler, error) {
	handler := &CloudWatchHandler{}

	switch os.Getenv("APP_ENV") {
	case "main":
		if err := handler.main(); err != nil {
			return nil, errors.NewMethodError("handler.main", err)
		}
	case "dev":
		if err := handler.dev(); err != nil {
			return nil, errors.NewMethodError("handler.dev", err)
		}
	default:
		return nil, errors.NewError("APP_ENV is invalid")
	}

	CloudWatchHandlerInstance = handler

	return CloudWatchHandlerInstance, nil
}

// dev 開発環境
func (s *CloudWatchHandler) dev() error {
	return nil
}

// main 本番環境
func (s *CloudWatchHandler) main() error {
	if err := s.userDB(); err != nil {
		return errors.NewMethodError("s.userDB", err)
	}

	return nil
}

// userDB コネクションを作成する
func (s *CloudWatchHandler) userDB() error {
	region := os.Getenv("AWS_REGION")

	readConn, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return errors.NewMethodError("config.LoadDefaultConfig", err)
	}

	writeConn, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return errors.NewMethodError("config.LoadDefaultConfig", err)
	}

	s.User = &CloudWatchConn{
		ReadMysqlConn:  cloudwatchlogs.NewFromConfig(readConn),
		WriteMysqlConn: cloudwatchlogs.NewFromConfig(writeConn),
	}

	if err := s.createUserLog("", ""); err != nil {
		return errors.NewMethodError("s.createUserLog", err)
	}

	return nil
}

// createUserLog ロググループとログストリームを作成する
func (s *CloudWatchHandler) createUserLog(logGroupName, logStreamName string) error {
	if _, err := s.User.WriteMysqlConn.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: &logGroupName,
	}); err != nil {
		return errors.NewMethodError("s.User.WriteMysqlConn.CreateLogGroup", err)
	}

	if _, err := s.User.WriteMysqlConn.CreateLogStream(context.TODO(), &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  &logGroupName,
		LogStreamName: &logStreamName,
	}); err != nil {
		return errors.NewMethodError("s.User.WriteMysqlConn.CreateLogStream", err)
	}

	return nil
}
