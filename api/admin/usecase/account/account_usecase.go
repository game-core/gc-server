package account

import (
	"context"

	accountProto "github.com/game-core/gc-server/api/admin/presentation/proto/account"
	adminAccountGoogleTokenProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
	adminAccountGoogleUrlProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleUrl"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	googleService "github.com/game-core/gc-server/pkg/domain/model/google"
)

type AccountUsecase interface {
	GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error)
	GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error)
}

type accountUsecase struct {
	googleService googleService.GoogleService
}

func NewAccountUsecase(
	googleService googleService.GoogleService,
) AccountUsecase {
	return &accountUsecase{
		googleService: googleService,
	}
}

// GetGoogleUrl URLを取得する
func (s *accountUsecase) GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error) {
	res, err := s.googleService.GetAdminGoogleUrl()
	if err != nil {
		return nil, errors.NewMethodError("s.googleService.GetAdminGoogleUrl", err)
	}

	return accountProto.SetAccountGetGoogleUrlResponse(
		adminAccountGoogleUrlProto.SetAdminAccountGoogleUrl(
			res.URL,
		),
	), nil
}

// GetGoogleToken トークンを取得する
func (s *accountUsecase) GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error) {
	res, err := s.googleService.GetAdminGoogleToken(ctx, req.Code)
	if err != nil {
		return nil, errors.NewMethodError("s.googleService.GetAdminGoogleToken", err)
	}

	return accountProto.SetAccountGetGoogleTokenResponse(
		adminAccountGoogleTokenProto.SetAdminAccountGoogleToken(
			res.AccessToken,
			res.RefreshToken,
			times.TimeToPb(&res.ExpiredAt),
		),
	), nil
}
