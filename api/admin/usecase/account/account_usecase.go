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
	GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error)
	GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error)
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

// GetGoogleUrl URLを取得する
func (s *accountUsecase) GetGoogleUrl(ctx context.Context, req *accountProto.AccountGetGoogleUrlRequest) (*accountProto.AccountGetGoogleUrlResponse, error) {
	return accountProto.SetAccountGetGoogleUrlResponse(""), nil
}

// GetGoogleToken トークンを取得する
func (s *accountUsecase) GetGoogleToken(ctx context.Context, req *accountProto.AccountGetGoogleTokenRequest) (*accountProto.AccountGetGoogleTokenResponse, error) {
	return accountProto.SetAccountGetGoogleTokenResponse(
		adminAccountGoogleTokenProto.SetAdminAccountGoogleToken(
			"",
			"",
			times.TimeToPb(&time.Time{}),
		),
	), nil
}
