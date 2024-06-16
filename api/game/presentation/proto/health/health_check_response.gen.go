// Package health ヘルスチェックレスポンス
package health

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/health/commonHealth"
	"github.com/game-core/gc-server/api/game/presentation/proto/health/masterHealth"
)

type HealthCheckResponses []*HealthCheckResponse

func NewHealthCheckResponse() *HealthCheckResponse {
	return &HealthCheckResponse{}
}

func NewHealthCheckResponses() HealthCheckResponses {
	return HealthCheckResponses{}
}

func SetHealthCheckResponse(commonHealth *commonHealth.CommonHealth, masterHealth *masterHealth.MasterHealth) *HealthCheckResponse {
	return &HealthCheckResponse{
		CommonHealth: commonHealth,
		MasterHealth: masterHealth,
	}
}
