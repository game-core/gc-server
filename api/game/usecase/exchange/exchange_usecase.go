package exchange

import (
	"context"

	exchangeProto "github.com/game-core/gc-server/api/game/presentation/proto/exchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange/userExchangeItem"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/times"
	exchangeService "github.com/game-core/gc-server/pkg/domain/model/exchange"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type ExchangeUsecase interface {
	Update(ctx context.Context, req *exchangeProto.ExchangeUpdateRequest) (*exchangeProto.ExchangeUpdateResponse, error)
	Receive(ctx context.Context, req *exchangeProto.ExchangeReceiveRequest) (*exchangeProto.ExchangeReceiveResponse, error)
}

type exchangeUsecase struct {
	exchangeService    exchangeService.ExchangeService
	transactionService transactionService.TransactionService
}

func NewExchangeUsecase(
	exchangeService exchangeService.ExchangeService,
	transactionService transactionService.TransactionService,
) ExchangeUsecase {
	return &exchangeUsecase{
		exchangeService:    exchangeService,
		transactionService: transactionService,
	}
}

// Update 交換情報を更新
func (s *exchangeUsecase) Update(ctx context.Context, req *exchangeProto.ExchangeUpdateRequest) (*exchangeProto.ExchangeUpdateResponse, error) {
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	res, err := s.exchangeService.Update(ctx, tx, times.Now(), exchangeService.SetExchangeUpdateRequest(req.UserId, req.MasterExchangeId))
	if err != nil {
		return nil, errors.NewMethodError("s.exchangeService.Update", err)
	}

	return exchangeProto.SetExchangeUpdateResponse(
		userExchange.SetUserExchange(
			res.UserExchange.UserId,
			res.UserExchange.MasterExchangeId,
			times.TimeToPb(&res.UserExchange.ResetAt),
		),
		userExchangeItem.SetUserExchangeItems(
			res.UserExchangeItems,
		),
	), nil
}

// Receive 交換受け取り
func (s *exchangeUsecase) Receive(ctx context.Context, req *exchangeProto.ExchangeReceiveRequest) (*exchangeProto.ExchangeReceiveResponse, error) {
	tx, err := s.transactionService.UserMysqlBegin(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.transactionService.UserMysqlBegin", err)
	}
	defer func() {
		s.transactionService.UserMysqlEnd(ctx, tx, err)
	}()

	res, err := s.exchangeService.Receive(ctx, tx, times.Now(), exchangeService.SetExchangeReceiveRequest(req.UserId, req.MasterExchangeItemId, req.Count))
	if err != nil {
		return nil, errors.NewMethodError("s.exchangeService.Receive", err)
	}

	return exchangeProto.SetExchangeReceiveResponse(
		userExchange.SetUserExchange(
			res.UserExchange.UserId,
			res.UserExchange.MasterExchangeId,
			times.TimeToPb(&res.UserExchange.ResetAt),
		),
		userExchangeItem.SetUserExchangeItem(
			res.UserExchangeItem.UserId,
			res.UserExchangeItem.MasterExchangeId,
			res.UserExchangeItem.MasterExchangeItemId,
			res.UserExchangeItem.Count,
		),
	), nil
}
