// Package adminAccountGoogleToken 管理者アカウントのGoogleToken
package adminAccountGoogleToken

import (
	"time"
)

type AdminAccountGoogleTokens []*AdminAccountGoogleToken

type AdminAccountGoogleToken struct {
	AccessToken  string
	RefreshToken string
	ExpiredAt    time.Time
}

func NewAdminAccountGoogleToken() *AdminAccountGoogleToken {
	return &AdminAccountGoogleToken{}
}

func NewAdminAccountGoogleTokens() AdminAccountGoogleTokens {
	return AdminAccountGoogleTokens{}
}

func SetAdminAccountGoogleToken(accessToken string, refreshToken string, expiredAt time.Time) *AdminAccountGoogleToken {
	return &AdminAccountGoogleToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
}
