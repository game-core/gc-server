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
	ReadCloudWatchConn  *cloudwatchlogs.Client
	WriteCloudWatchConn *cloudwatchlogs.Client
}

// NewCloudWatch インスタンスを作成する
func NewCloudWatch() *CloudWatchHandler {
	return CloudWatchHandlerInstance
}

// InitCloudWatch 初期化する
func InitCloudWatch() (*CloudWatchHandler, error) {
	handler := &CloudWatchHandler{}

	switch os.Getenv("APP_ENV") {
	case "prod":
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
	logGroupName := os.Getenv("USER_LOG_GROUP_NAME")
	logStreamName := os.Getenv("USER_LOG_STREAM_NAME")

	readConn, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return errors.NewMethodError("config.LoadDefaultConfig", err)
	}

	writeConn, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return errors.NewMethodError("config.LoadDefaultConfig", err)
	}

	s.User = &CloudWatchConn{
		ReadCloudWatchConn:  cloudwatchlogs.NewFromConfig(readConn),
		WriteCloudWatchConn: cloudwatchlogs.NewFromConfig(writeConn),
	}

	if err := s.createUserLog(logGroupName, logStreamName); err != nil {
		return errors.NewMethodError("s.createUserLog", err)
	}

	return nil
}

// createUserLog ロググループとログストリームを作成する
func (s *CloudWatchHandler) createUserLog(logGroupName, logStreamName string) error {
	if err := s.createUserLogGroup(logGroupName); err != nil {
		return errors.NewMethodError("s.createUserLogGroup", err)
	}

	if err := s.createUserLogStream(logGroupName, logStreamName); err != nil {
		return errors.NewMethodError("s.createUserLogStream", err)
	}

	return nil
}

// createUserLogGroup ロググループを作成
func (s *CloudWatchHandler) createUserLogGroup(logGroupName string) error {
	logGroupExists, err := s.checkUserLogGroupExist(logGroupName)
	if err != nil {
		return errors.NewMethodError("s.checkUserLogGroupExist", err)
	}
	if !logGroupExists {
		if _, err := s.User.WriteCloudWatchConn.CreateLogGroup(context.TODO(), &cloudwatchlogs.CreateLogGroupInput{
			LogGroupName: &logGroupName,
		}); err != nil {
			return errors.NewMethodError("s.User.WriteCloudWatchConn.CreateLogGroup", err)
		}
	}

	return nil
}

// createUserLogStream ログストリームを作成
func (s *CloudWatchHandler) createUserLogStream(logGroupName, logStreamName string) error {
	logStreamExists, err := s.checkUserLogStreamExist(logGroupName, logStreamName)
	if err != nil {
		return errors.NewMethodError("s.checkUserLogStreamExist", err)
	}
	if !logStreamExists {
		if _, err := s.User.WriteCloudWatchConn.CreateLogStream(context.TODO(), &cloudwatchlogs.CreateLogStreamInput{
			LogGroupName:  &logGroupName,
			LogStreamName: &logStreamName,
		}); err != nil {
			return errors.NewMethodError("s.User.WriteCloudWatchConn.CreateLogStream", err)
		}
	}

	return nil
}

// checkUserLogGroupExist ロググループが存在するか確認する
func (s *CloudWatchHandler) checkUserLogGroupExist(logGroupName string) (bool, error) {
	res, err := s.User.ReadCloudWatchConn.DescribeLogGroups(context.TODO(), &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: &logGroupName,
	})
	if err != nil {
		return false, err
	}

	for _, logGroup := range res.LogGroups {
		if *logGroup.LogGroupName == logGroupName {
			return true, nil
		}
	}

	return false, nil
}

// checkUserLogStreamExist ログストリームが存在するか確認する
func (s *CloudWatchHandler) checkUserLogStreamExist(logGroupName, logStreamName string) (bool, error) {
	res, err := s.User.ReadCloudWatchConn.DescribeLogStreams(context.TODO(), &cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        &logGroupName,
		LogStreamNamePrefix: &logStreamName,
	})
	if err != nil {
		return false, err
	}

	for _, logStream := range res.LogStreams {
		if *logStream.LogStreamName == logStreamName {
			return true, nil
		}
	}

	return false, nil
}
