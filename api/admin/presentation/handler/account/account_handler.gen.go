package account

import (
	"context"

	"github.com/game-core/gc-server/api/admin/presentation/proto/account"
	accountUsecase "github.com/game-core/gc-server/api/admin/usecase/account"
	"github.com/game-core/gc-server/internal/errors"
)

type AccountHandler interface {
	account.AccountServer
}

type accountHandler struct {
	account.UnimplementedAccountServer
	accountUsecase accountUsecase.AccountUsecase
}

func NewAccountHandler(
	accountUsecase accountUsecase.AccountUsecase,
) AccountHandler {
	return &accountHandler{
		accountUsecase: accountUsecase,
	}
}

// GetGoogleUrl GoogleURLを取得する
func (s *accountHandler) GetGoogleUrl(ctx context.Context, req *account.AccountGetGoogleUrlRequest) (*account.AccountGetGoogleUrlResponse, error) {
	res, err := s.accountUsecase.GetGoogleUrl(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.GetGoogleUrl", err)
	}

	return res, nil
}

// GetGoogleToken GoogleTokenを取得する
func (s *accountHandler) GetGoogleToken(ctx context.Context, req *account.AccountGetGoogleTokenRequest) (*account.AccountGetGoogleTokenResponse, error) {
	res, err := s.accountUsecase.GetGoogleToken(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.GetGoogleToken", err)
	}

	return res, nil
}
