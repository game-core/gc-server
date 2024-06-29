package adminGoogle

import "time"

type AdminGoogleToken struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
}

func SetAdminGoogleToken(accessToken, refreshToken string, expiredAt time.Time) *AdminGoogleToken {
	return &AdminGoogleToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
}
