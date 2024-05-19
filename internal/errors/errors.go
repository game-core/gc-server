package errors

import (
	"fmt"
	"log"
)

// NewError エラー
func NewError(message string) error {
	return fmt.Errorf(message)
}

// NewMethodError エラー
func NewMethodError(method string, err error) error {
	return fmt.Errorf("failed to %s: %s", method, err)
}

// NewErrorLog エラーログ
func NewErrorLog(message string) {
	log.Printf(message)
}

// NewMethodErrorLog エラーログ
func NewMethodErrorLog(method string, err error) {
	log.Printf("failed to %s: %s", method, err)
}

// NewTestError テストエラー
func NewTestError() error {
	return fmt.Errorf("test")
}
