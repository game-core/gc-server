// Package health ヘルスチェックリクエスト
package health

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/health/commonHealth"
	"github.com/game-core/gc-server/api/game/presentation/proto/health/masterHealth"
)

type HealthCheckRequests []*HealthCheckRequest

func NewHealthCheckRequest() *HealthCheckRequest {
	return &HealthCheckRequest{}
}

func NewHealthCheckRequests() HealthCheckRequests {
	return HealthCheckRequests{}
}

func SetHealthCheckRequest(healthId int64, name string, commonHealthType commonHealth.CommonHealthType, masterHealthType masterHealth.MasterHealthType) *HealthCheckRequest {
	return &HealthCheckRequest{
		HealthId:         healthId,
		Name:             name,
		CommonHealthType: commonHealthType,
		MasterHealthType: masterHealthType,
	}
}
