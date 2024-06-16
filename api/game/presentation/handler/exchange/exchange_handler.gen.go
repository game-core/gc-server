package exchange

import (
	"context"

	"github.com/game-core/gc-server/api/game/presentation/proto/exchange"
	exchangeUsecase "github.com/game-core/gc-server/api/game/usecase/exchange"
	"github.com/game-core/gc-server/internal/errors"
)

type ExchangeHandler interface {
	exchange.ExchangeServer
}

type exchangeHandler struct {
	exchange.UnimplementedExchangeServer
	exchangeUsecase exchangeUsecase.ExchangeUsecase
}

func NewExchangeHandler(
	exchangeUsecase exchangeUsecase.ExchangeUsecase,
) ExchangeHandler {
	return &exchangeHandler{
		exchangeUsecase: exchangeUsecase,
	}
}

// Update 更新
func (s *exchangeHandler) Update(ctx context.Context, req *exchange.ExchangeUpdateRequest) (*exchange.ExchangeUpdateResponse, error) {
	res, err := s.exchangeUsecase.Update(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.exchangeUsecase.Update", err)
	}

	return res, nil
}

// Receive 受け取り
func (s *exchangeHandler) Receive(ctx context.Context, req *exchange.ExchangeReceiveRequest) (*exchange.ExchangeReceiveResponse, error) {
	res, err := s.exchangeUsecase.Receive(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.exchangeUsecase.Receive", err)
	}

	return res, nil
}
