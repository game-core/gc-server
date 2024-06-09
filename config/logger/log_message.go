package logger

import (
	"encoding/json"
	"time"

	"github.com/game-core/gc-server/internal/errors"
)

type LogMessage struct {
	Timestamp time.Time   `json:"timestamp"`
	LogLevel  LogLevel    `json:"level"`
	Message   interface{} `json:"message"`
}

func SetLogMessage(timestamp time.Time, logLevel LogLevel, message interface{}) *LogMessage {
	return &LogMessage{
		Timestamp: timestamp,
		LogLevel:  logLevel,
		Message:   message,
	}
}

func (t *LogMessage) ToJson() []byte {
	messageJSON, err := json.Marshal(t)
	if err != nil {
		errors.NewMethodErrorLog("json.Marshal", err)
		return nil
	}

	return messageJSON
}
