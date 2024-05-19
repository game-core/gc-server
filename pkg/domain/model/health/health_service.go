//go:generate mockgen -source=./health_service.go -destination=./health_service_mock.gen.go -package=health
package health

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type HealthService interface {
	Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error)
}

type healthService struct {
	commonHealthMysqlRepository commonHealth.CommonHealthMysqlRepository
	masterHealthMysqlRepository masterHealth.MasterHealthMysqlRepository
}

func NewHealthService(
	commonHealthMysqlRepository commonHealth.CommonHealthMysqlRepository,
	masterHealthMysqlRepository masterHealth.MasterHealthMysqlRepository,
) HealthService {
	return &healthService{
		commonHealthMysqlRepository: commonHealthMysqlRepository,
		masterHealthMysqlRepository: masterHealthMysqlRepository,
	}
}

// Check ヘルスチェック
func (s *healthService) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	commonHealthModel, err := s.commonHealthMysqlRepository.Find(ctx, req.HealthId)
	if err != nil {
		return nil, errors.NewMethodError("s.commonHealthMysqlRepository.Find", err)
	}

	masterHealthModel, err := s.masterHealthMysqlRepository.Find(ctx, req.HealthId)
	if err != nil {
		return nil, errors.NewMethodError("s.masterHealthMysqlRepository.Find", err)
	}

	return SetHealthCheckResponse(
		commonHealthModel,
		masterHealthModel,
	), nil
}
