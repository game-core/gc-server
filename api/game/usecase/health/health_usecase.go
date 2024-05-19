package health

import (
	"context"

	healthServer "github.com/game-core/gc-server/api/game/presentation/server/health"
	"github.com/game-core/gc-server/api/game/presentation/server/health/commonHealth"
	"github.com/game-core/gc-server/api/game/presentation/server/health/masterHealth"
	"github.com/game-core/gc-server/internal/errors"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
	commonHealthModel "github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	masterHealthModel "github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type HealthUsecase interface {
	Check(ctx context.Context, req *healthServer.HealthCheckRequest) (*healthServer.HealthCheckResponse, error)
}

type healthUsecase struct {
	healthService healthService.HealthService
}

func NewHealthUsecase(
	healthService healthService.HealthService,
) HealthUsecase {
	return &healthUsecase{
		healthService: healthService,
	}
}

// Check ヘルスチェック
func (s *healthUsecase) Check(ctx context.Context, req *healthServer.HealthCheckRequest) (*healthServer.HealthCheckResponse, error) {
	res, err := s.healthService.Check(
		ctx,
		healthService.SetHealthCheckRequest(
			req.HealthId,
			req.Name,
			commonHealthModel.CommonHealthType(req.CommonHealthType),
			masterHealthModel.MasterHealthType(req.MasterHealthType),
		),
	)
	if err != nil {
		return nil, errors.NewMethodError("s.healthService.Check", err)
	}

	return healthServer.SetHealthCheckResponse(
		commonHealth.SetCommonHealth(
			res.CommonHealth.HealthId,
			res.CommonHealth.Name,
			commonHealth.CommonHealthType(res.CommonHealth.CommonHealthType),
		),
		masterHealth.SetMasterHealth(
			res.MasterHealth.HealthId,
			res.MasterHealth.Name,
			masterHealth.MasterHealthType(res.MasterHealth.MasterHealthType),
		),
	), nil
}
