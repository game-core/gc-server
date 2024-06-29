package account

import (
	"context"

	accountProto "github.com/game-core/gc-server/api/admin/presentation/proto/account"
	adminAccountGoogleTokenProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
	adminAccountGoogleUrlProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleUrl"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
)

type AccountUsecase interface {
	GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error)
	GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error)
}

type accountUsecase struct {
	accountService accountService.AccountService
}

func NewAccountUsecase(
	accountService accountService.AccountService,
) AccountUsecase {
	return &accountUsecase{
		accountService: accountService,
	}
}

// GetGoogleUrl URLを取得する
func (s *accountUsecase) GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error) {
	res, err := s.accountService.GetGoogleUrl()
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.GetGoogleUrl", err)
	}

	return accountProto.SetAccountGetGoogleUrlResponse(
		adminAccountGoogleUrlProto.SetAdminAccountGoogleUrl(
			res.AdminAccountGoogleUrl.Url,
		),
	), nil
}

// GetGoogleToken トークンを取得する
func (s *accountUsecase) GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error) {
	res, err := s.accountService.GetGoogleToken(ctx, accountService.SetAccountGetGoogleTokenRequest(req.Code))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.GetGoogleUrl", err)
	}

	return accountProto.SetAccountGetGoogleTokenResponse(
		adminAccountGoogleTokenProto.SetAdminAccountGoogleToken(
			res.AdminAccountGoogleToken.AccessToken,
			res.AdminAccountGoogleToken.RefreshToken,
			times.TimeToPb(&res.AdminAccountGoogleToken.ExpiredAt),
		),
	), nil
}
