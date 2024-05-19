// Package health ヘルスチェックレスポンス
package health

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type HealthCheckResponses []*HealthCheckResponse

type HealthCheckResponse struct {
	CommonHealth *commonHealth.CommonHealth
	MasterHealth *masterHealth.MasterHealth
}

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
