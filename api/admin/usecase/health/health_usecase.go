package health

import (
	"context"

	healthProto "github.com/game-core/gc-server/api/admin/presentation/proto/health"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/adminHealth"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/commonHealth"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/masterHealth"
	"github.com/game-core/gc-server/internal/errors"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
)

type HealthUsecase interface {
	Check(ctx context.Context, req *healthProto.HealthCheckRequest) (*healthProto.HealthCheckResponse, error)
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
func (s *healthUsecase) Check(ctx context.Context, req *healthProto.HealthCheckRequest) (*healthProto.HealthCheckResponse, error) {
	res, err := s.healthService.Check(
		ctx,
		healthService.SetHealthCheckRequest(
			req.HealthId,
		),
	)
	if err != nil {
		return nil, errors.NewMethodError("s.healthService.Check", err)
	}

	return healthProto.SetHealthCheckResponse(
		adminHealth.SetAdminHealth(
			res.AdminHealth.HealthId,
			res.AdminHealth.Name,
			adminHealth.AdminHealthEnum(res.AdminHealth.AdminHealthEnum),
		),
		commonHealth.SetCommonHealth(
			res.CommonHealth.HealthId,
			res.CommonHealth.Name,
			commonHealth.CommonHealthEnum(res.CommonHealth.CommonHealthEnum),
		),
		masterHealth.SetMasterHealth(
			res.MasterHealth.HealthId,
			res.MasterHealth.Name,
			masterHealth.MasterHealthEnum(res.MasterHealth.MasterHealthEnum),
		),
	), nil
}
