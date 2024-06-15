package userLoginBonus

import (
	"time"
)

// CheckReceived 報酬を受け取っているか確認する(true=受け取り済, false=受け取り無)
func (s *UserLoginBonus) CheckReceived(resetHour int32, intervalHour int32, now time.Time) bool {
	if s != nil {
		resetTime := time.Date(now.Year(), now.Month(), now.Day(), int(resetHour), 0, 0, 0, now.Location())
		if now.Before(resetTime) {
			return !s.ReceivedAt.Add(time.Duration(intervalHour) * time.Hour).Before(now)
		}

		return !s.ReceivedAt.Before(resetTime)
	}

	return false
}
