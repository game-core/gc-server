package health

import (
	"context"

	"github.com/game-core/gc-server/api/game/presentation/proto/health"
	healthUsecase "github.com/game-core/gc-server/api/game/usecase/health"
	"github.com/game-core/gc-server/internal/errors"
)

type HealthHandler interface {
	health.HealthServer
}

type healthHandler struct {
	health.UnimplementedHealthServer
	healthUsecase healthUsecase.HealthUsecase
}

func NewHealthHandler(
	healthUsecase healthUsecase.HealthUsecase,
) HealthHandler {
	return &healthHandler{
		healthUsecase: healthUsecase,
	}
}

// Check ヘルスチェック
func (s *healthHandler) Check(ctx context.Context, req *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	res, err := s.healthUsecase.Check(ctx, req)
	if err != nil {
		return nil, errors.NewMethodError("s.healthUsecase.Check", err)
	}

	return res, nil
}
