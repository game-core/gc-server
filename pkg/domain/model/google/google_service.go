//go:generate mockgen -source=./google_service.go -destination=./google_service_mock.gen.go -package=google
package google

import (
	"context"

	"github.com/game-core/gc-server/pkg/domain/model/google/adminGoogle"
)

type GoogleService interface {
	GetAdminGoogleUrl() (*adminGoogle.AdminGoogleURL, error)
	GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error)
	GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error)
}

type googleService struct {
	adminGoogleAuthRepository adminGoogle.AdminGoogleAuthRepository
}

func NewGoogleService(
	adminGoogleAuthRepository adminGoogle.AdminGoogleAuthRepository,
) GoogleService {
	return &googleService{
		adminGoogleAuthRepository: adminGoogleAuthRepository,
	}
}

// GetAdminGoogleUrl URLを取得する
func (s *googleService) GetAdminGoogleUrl() (*adminGoogle.AdminGoogleURL, error) {
	return s.adminGoogleAuthRepository.GetAdminGoogleUrl()
}

// GetAdminGoogleToken Tokenを取得する
func (s *googleService) GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error) {
	return s.adminGoogleAuthRepository.GetAdminGoogleToken(ctx, code)
}

// GetAdminGoogleTokenInfo Token情報を取得する
func (s *googleService) GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error) {
	return s.adminGoogleAuthRepository.GetAdminGoogleTokenInfo(ctx, accessToken)
}
