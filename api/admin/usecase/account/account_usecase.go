package account

import (
	"context"
	"time"

	accountProto "github.com/game-core/gc-server/api/admin/presentation/proto/account"
	adminAccountGoogleTokenProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
	"github.com/game-core/gc-server/internal/times"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type AccountUsecase interface {
	GetGoogleLoginUrl(ctx context.Context, req *accountProto.AccountGetGoogleLoginUrlRequest) (*accountProto.AccountGetGoogleLoginUrlResponse, error)
	GetGoogleLoginToken(ctx context.Context, req *accountProto.AccountGetGoogleLoginTokenRequest) (*accountProto.AccountGetGoogleLoginTokenResponse, error)
}

type accountUsecase struct {
	accountService     accountService.AccountService
	transactionService transactionService.TransactionService
}

func NewAccountUsecase(
	accountService accountService.AccountService,
	transactionService transactionService.TransactionService,
) AccountUsecase {
	return &accountUsecase{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

// GetGoogleLoginUrl アカウントをログインする
func (s *accountUsecase) GetGoogleLoginUrl(ctx context.Context, req *accountProto.AccountGetGoogleLoginUrlRequest) (*accountProto.AccountGetGoogleLoginUrlResponse, error) {
	return accountProto.SetAccountGetGoogleLoginUrlResponse(""), nil
}

func (s *accountUsecase) GetGoogleLoginToken(ctx context.Context, req *accountProto.AccountGetGoogleLoginTokenRequest) (*accountProto.AccountGetGoogleLoginTokenResponse, error) {
	return accountProto.SetAccountGetGoogleLoginTokenResponse(
		adminAccountGoogleTokenProto.SetAdminAccountGoogleToken(
			"",
			"",
			times.TimeToPb(&time.Time{}),
		),
	), nil
}
