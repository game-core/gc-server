// Package adminGoogle 管理者用Google
//
//go:generate mockgen -source=./admin_google_auth_repository.go -destination=./admin_google_auth_repository_mock.gen.go -package=adminGoogle
package adminGoogle

import "context"

type AdminGoogleAuthRepository interface {
	GetAdminGoogleUrl() (*AdminGoogleURL, error)
	GetAdminGoogleToken(ctx context.Context, code string) (*AdminGoogleToken, error)
	GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*AdminGoogleTokenInfo, error)
	RefreshAdminGoogleToken(ctx context.Context, refreshToken string) (*AdminGoogleToken, error)
}
