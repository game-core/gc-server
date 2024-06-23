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

// GetGoogleLoginUrl GoogleログインURLを取得する
func (s *accountHandler) GetGoogleLoginUrl(ctx context.Context, req *account.AccountGetGoogleLoginUrlRequest) (*account.AccountGetGoogleLoginUrlResponse, error) {
	res, err := s.accountUsecase.GetGoogleLoginUrl(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.GetGoogleLoginUrl", err)
	}

	return res, nil
}

// GetGoogleLoginToken GoogleログインTokenを取得する
func (s *accountHandler) GetGoogleLoginToken(ctx context.Context, req *account.AccountGetGoogleLoginTokenRequest) (*account.AccountGetGoogleLoginTokenResponse, error) {
	res, err := s.accountUsecase.GetGoogleLoginToken(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.accountUsecase.GetGoogleLoginToken", err)
	}

	return res, nil
}
