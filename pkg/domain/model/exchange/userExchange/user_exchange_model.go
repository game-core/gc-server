package userExchange

import (
	"time"
)

// CheckResetAt リセットが必要か確認する(true=リセット有, false=リセット無)
func (s *UserExchange) CheckResetAt(now time.Time, intervalHour int32) bool {
	return now.Sub(s.ResetAt).Hours() >= float64(intervalHour)
}

// CreateResetAt 新しいリセット日時を作成する
func (s *UserExchange) CreateResetAt(now, startAt time.Time, resetHour, intervalHour int32) time.Time {
	if s == nil {
		initialResetAt := time.Date(startAt.Year(), startAt.Month(), startAt.Day(), int(resetHour), 0, 0, 0, startAt.Location())
		if startAt.After(initialResetAt) {
			initialResetAt = initialResetAt.Add(time.Duration(intervalHour) * time.Hour)
		}
		for initialResetAt.Add(time.Duration(intervalHour) * time.Hour).Before(now) {
			initialResetAt = initialResetAt.Add(time.Duration(intervalHour) * time.Hour)
		}
		return initialResetAt
	}

	return s.ResetAt.Add(time.Duration(intervalHour) * time.Hour)
}
