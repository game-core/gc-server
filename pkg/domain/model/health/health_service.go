//go:generate mockgen -source=./health_service.go -destination=./health_service_mock.gen.go -package=health
package health

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/adminHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type HealthService interface {
	Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error)
}

type healthService struct {
	adminHealthMysqlRepository  adminHealth.AdminHealthMysqlRepository
	commonHealthMysqlRepository commonHealth.CommonHealthMysqlRepository
	masterHealthMysqlRepository masterHealth.MasterHealthMysqlRepository
}

func NewHealthService(
	adminHealthMysqlRepository adminHealth.AdminHealthMysqlRepository,
	commonHealthMysqlRepository commonHealth.CommonHealthMysqlRepository,
	masterHealthMysqlRepository masterHealth.MasterHealthMysqlRepository,
) HealthService {
	return &healthService{
		adminHealthMysqlRepository:  adminHealthMysqlRepository,
		commonHealthMysqlRepository: commonHealthMysqlRepository,
		masterHealthMysqlRepository: masterHealthMysqlRepository,
	}
}

// Check ヘルスチェック
func (s *healthService) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	adminHealthModel, err := s.adminHealthMysqlRepository.Find(ctx, req.HealthId)
	if err != nil {
		return nil, errors.NewMethodError("s.adminHealthMysqlRepository.Find", err)
	}

	commonHealthModel, err := s.commonHealthMysqlRepository.Find(ctx, req.HealthId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonHealthMysqlRepository.Find", err)
	}

	masterHealthModel, err := s.masterHealthMysqlRepository.Find(ctx, req.HealthId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterHealthMysqlRepository.Find", err)
	}

	return SetHealthCheckResponse(
		adminHealthModel,
		commonHealthModel,
		masterHealthModel,
	), nil
}
