package loginBonus

import (
	"context"

	"github.com/game-core/gc-server/api/game/presentation/server/loginBonus"
	loginBonusUsecase "github.com/game-core/gc-server/api/game/usecase/loginBonus"
	"github.com/game-core/gc-server/internal/errors"
)

type LoginBonusHandler interface {
	loginBonus.LoginBonusServer
}

type loginBonusHandler struct {
	loginBonus.UnimplementedLoginBonusServer
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase
}

func NewLoginBonusHandler(
	loginBonusUsecase loginBonusUsecase.LoginBonusUsecase,
) LoginBonusHandler {
	return &loginBonusHandler{
		loginBonusUsecase: loginBonusUsecase,
	}
}

// Receive 受け取り
func (s *loginBonusHandler) Receive(ctx context.Context, req *loginBonus.LoginBonusReceiveRequest) (*loginBonus.LoginBonusReceiveResponse, error) {
	res, err := s.loginBonusUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.loginBonusUsecase.Receive", err)
	}

	return res, nil
}
