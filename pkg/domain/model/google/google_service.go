//go:generate mockgen -source=./google_service.go -destination=./google_service_mock.gen.go -package=google
package google

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/google/adminGoogle"
)

type GoogleService interface {
	GetAdminGoogleUrl() (*adminGoogle.AdminGoogleURL, error)
	GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error)
	GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error)
	RefreshAdminGoogleToken(ctx context.Context, refreshToken string) (*adminGoogle.AdminGoogleToken, error)
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
	adminGoogleURL, err := s.adminGoogleAuthRepository.GetAdminGoogleUrl()
	if err != nil {
		return nil, errors.NewMethodError("s.adminGoogleAuthRepository.GetAdminGoogleUrl", err)
	}

	return adminGoogleURL, nil
}

// GetAdminGoogleToken Tokenを取得する
func (s *googleService) GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error) {
	adminGoogleURL, err := s.adminGoogleAuthRepository.GetAdminGoogleToken(ctx, code)
	if err != nil {
		return nil, errors.NewMethodError("s.adminGoogleAuthRepository.GetAdminGoogleToken", err)
	}

	return adminGoogleURL, nil
}

// GetAdminGoogleTokenInfo Token情報を取得する
func (s *googleService) GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error) {
	adminGoogleURL, err := s.adminGoogleAuthRepository.GetAdminGoogleTokenInfo(ctx, accessToken)
	if err != nil {
		return nil, errors.NewMethodError("s.adminGoogleAuthRepository.GetAdminGoogleTokenInfo", err)
	}

	return adminGoogleURL, nil
}

// RefreshAdminGoogleToken Tokenをリフレッシュする
func (s *googleService) RefreshAdminGoogleToken(ctx context.Context, refreshToken string) (*adminGoogle.AdminGoogleToken, error) {
	adminGoogleURL, err := s.adminGoogleAuthRepository.RefreshAdminGoogleToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.NewMethodError("s.adminGoogleAuthRepository.RefreshAdminGoogleToken", err)
	}

	return adminGoogleURL, nil
}
