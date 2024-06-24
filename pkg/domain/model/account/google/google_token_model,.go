package google

import "time"

type GoogleToken struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
}

func SetGoogleToken(accessToken, refreshToken string, expiredAt time.Time) *GoogleToken {
	return &GoogleToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
}
