package userLoginBonus

import (
	"time"
)

// CheckReceived 報酬を受け取っているか確認する(true=受け取り済, false=受け取り無)
func (s *UserLoginBonus) CheckReceived(resetHour int32, intervalHour int32, now time.Time) bool {
	if s != nil {
		resetDuration := time.Duration(intervalHour) * time.Hour
		elapsed := now.Sub(time.Date(now.Year(), now.Month(), now.Day(), int(resetHour), 0, 0, 0, now.Location()))
		lastResetTime := now.Add(-elapsed).Truncate(resetDuration).Add(time.Duration(resetHour) * time.Hour)

		if now.Before(lastResetTime) {
			lastResetTime = lastResetTime.Add(-resetDuration)
		}

		return s.ReceivedAt.After(lastResetTime)
	}

	return false
}
