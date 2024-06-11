package loginBonus

import (
	"context"

	loginBonusProto "github.com/game-core/gc-server/api/game/presentation/proto/loginBonus"
	"github.com/game-core/gc-server/api/game/presentation/proto/loginBonus/userLoginBonus"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	loginBonusService "github.com/game-core/gc-server/pkg/domain/model/loginBonus"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type LoginBonusUsecase interface {
	Receive(ctx context.Context, req *loginBonusProto.LoginBonusReceiveRequest) (*loginBonusProto.LoginBonusReceiveResponse, error)
}

type loginBonusUsecase struct {
	loginBonusService  loginBonusService.LoginBonusService
	transactionService transactionService.TransactionService
}

func NewLoginBonusUsecase(
	loginBonusService loginBonusService.LoginBonusService,
	transactionService transactionService.TransactionService,
) LoginBonusUsecase {
	return &loginBonusUsecase{
		loginBonusService:  loginBonusService,
		transactionService: transactionService,
	}
}

// Receive ログインボーナス受け取り
func (s *loginBonusUsecase) Receive(ctx context.Context, req *loginBonusProto.LoginBonusReceiveRequest) (*loginBonusProto.LoginBonusReceiveResponse, error) {
	// transaction
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	res, err := s.loginBonusService.Receive(ctx, tx, times.Now(), loginBonusService.SetLoginBonusReceiveRequest(req.UserId, req.MasterLoginBonusId))
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusService.Receive", err)
	}

	return loginBonusProto.SetLoginBonusReceiveResponse(
		userLoginBonus.SetUserLoginBonus(
			res.UserLoginBonus.UserId,
			res.UserLoginBonus.MasterLoginBonusId,
			times.TimeToPb(&res.UserLoginBonus.ReceivedAt),
		),
	), nil
}
