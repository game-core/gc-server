// Package adminAccountGoogleToken 管理者アカウントのGoogleToken
package adminAccountGoogleToken

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AdminAccountGoogleTokens []*AdminAccountGoogleToken

func NewAdminAccountGoogleToken() *AdminAccountGoogleToken {
	return &AdminAccountGoogleToken{}
}

func NewAdminAccountGoogleTokens() AdminAccountGoogleTokens {
	return AdminAccountGoogleTokens{}
}

func SetAdminAccountGoogleToken(accessToken string, refreshToken string, expiredAt *timestamppb.Timestamp) *AdminAccountGoogleToken {
	return &AdminAccountGoogleToken{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiredAt:    expiredAt,
	}
}
