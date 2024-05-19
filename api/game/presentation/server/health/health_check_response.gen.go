// Package health ヘルスチェックレスポンス
package health

import (
	"github.com/game-core/gc-server/api/game/presentation/server/health/commonHealth"
	"github.com/game-core/gc-server/api/game/presentation/server/health/masterHealth"
)

func SetHealthCheckResponse(commonHealth *commonHealth.CommonHealth, masterHealth *masterHealth.MasterHealth) *HealthCheckResponse {
	return &HealthCheckResponse{
		CommonHealth: commonHealth,
		MasterHealth: masterHealth,
	}
}
