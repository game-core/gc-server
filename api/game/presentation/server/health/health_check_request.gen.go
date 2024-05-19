// Package health ヘルスチェックリクエスト
package health

import (
	"github.com/game-core/gc-server/api/game/presentation/server/health/commonHealth"
	"github.com/game-core/gc-server/api/game/presentation/server/health/masterHealth"
)

func SetHealthCheckRequest(healthId int64, name string, commonHealthType commonHealth.CommonHealthType, masterHealthType masterHealth.MasterHealthType) *HealthCheckRequest {
	return &HealthCheckRequest{
		HealthId:         healthId,
		Name:             name,
		CommonHealthType: commonHealthType,
		MasterHealthType: masterHealthType,
	}
}
