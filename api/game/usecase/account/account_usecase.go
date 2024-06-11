package account

import (
	"context"

	accountProto "github.com/game-core/gc-server/api/game/presentation/proto/account"
	"github.com/game-core/gc-server/api/game/presentation/proto/account/userAccount"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type AccountUsecase interface {
	Get(ctx context.Context, req *accountProto.AccountGetRequest) (*accountProto.AccountGetResponse, error)
	Create(ctx context.Context, req *accountProto.AccountCreateRequest) (*accountProto.AccountCreateResponse, error)
	Login(ctx context.Context, req *accountProto.AccountLoginRequest) (*accountProto.AccountLoginResponse, error)
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

// Get アカウントを確認する
func (s *accountUsecase) Get(ctx context.Context, req *accountProto.AccountGetRequest) (*accountProto.AccountGetResponse, error) {
	result, err := s.accountService.Get(ctx, accountService.SetAccountGetRequest(req.UserId))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Get", err)
	}

	return accountProto.SetAccountGetResponse(
		userAccount.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}

// Create アカウントを作成する
func (s *accountUsecase) Create(ctx context.Context, req *accountProto.AccountCreateRequest) (*accountProto.AccountCreateResponse, error) {
	userId, err := s.accountService.CreateUserId(ctx)
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.CreateUserId", err)
	}

	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, userId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	result, err := s.accountService.Create(ctx, tx, accountService.SetAccountCreateRequest(userId, req.Name, req.Password))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Create", err)
	}

	return accountProto.SetAccountCreateResponse(
		userAccount.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}

// Login アカウントをログインする
func (s *accountUsecase) Login(ctx context.Context, req *accountProto.AccountLoginRequest) (*accountProto.AccountLoginResponse, error) {
	// transaction
	mtx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	rtx := s.transactionService.UserRedisBegin()
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, mtx, err)
		s.transactionService.UserRedisEnd(ctx, rtx, err)
	}()

	result, err := s.accountService.Login(ctx, mtx, rtx, accountService.SetAccountLoginRequest(req.UserId, req.Password))
	if err != nil {
		return nil, errors.NewMethodError("s.accountService.Login", err)
	}

	return accountProto.SetAccountLoginResponse(
		result.Token,
		userAccount.SetUserAccount(
			result.UserAccount.UserId,
			result.UserAccount.Name,
			result.UserAccount.Password,
			times.TimeToPb(&result.UserAccount.LoginAt),
			times.TimeToPb(&result.UserAccount.LogoutAt),
		),
	), nil
}
