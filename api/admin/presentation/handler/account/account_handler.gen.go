package account

import (
	"context"

	"github.com/game-core/gc-server/api/admin/presentation/proto/account"
	accountUsecase "github.com/game-core/gc-server/api/admin/usecase/account"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/tokens"
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

// Get アカウントを取得する
		func (s *accountHandler) Get(ctx context.Context, req *account.AccountGetRequest) (*account.AccountGetResponse, error) {if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
				return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
			}
			res, err := s.accountUsecase.Get(ctx, req)
			if err != nil {
				return nil, errors.NewMethodError("s.accountUsecase.Get", err)
			}
		
			return res, nil
		}
		
// Create アカウントを作成する
		func (s *accountHandler) Create(ctx context.Context, req *account.AccountCreateRequest) (*account.AccountCreateResponse, error) {
			res, err := s.accountUsecase.Create(ctx, req)
			if err != nil {
				return nil, errors.NewMethodError("s.accountUsecase.Create", err)
			}
		
			return res, nil
		}
		
// Login アカウントをログインする
		func (s *accountHandler) Login(ctx context.Context, req *account.AccountLoginRequest) (*account.AccountLoginResponse, error) {
			res, err := s.accountUsecase.Login(ctx, req)
			if err != nil {
				return nil, errors.NewMethodError("s.accountUsecase.Login", err)
			}
		
			return res, nil
		}
		
