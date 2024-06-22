
// Package health ヘルスチェックレスポンス
package health

import (
	
"github.com/game-core/gc-server/api/game/presentation/proto/health/adminHealth"
"github.com/game-core/gc-server/api/game/presentation/proto/health/adminHealth"
"github.com/game-core/gc-server/api/game/presentation/proto/health/commonHealth"
"github.com/game-core/gc-server/api/game/presentation/proto/health/commonHealth"
"github.com/game-core/gc-server/api/game/presentation/proto/health/masterHealth"
"github.com/game-core/gc-server/api/game/presentation/proto/health/masterHealth"
)

type HealthCheckResponses []*HealthCheckResponse

func NewHealthCheckResponse() *HealthCheckResponse {
			return &HealthCheckResponse{}
		}

		func NewHealthCheckResponses() HealthCheckResponses {
			return HealthCheckResponses{}
		}

		func SetHealthCheckResponse(adminHealth *adminHealth.AdminHealth,commonHealth *commonHealth.CommonHealth,masterHealth *masterHealth.MasterHealth) *HealthCheckResponse {
			return &HealthCheckResponse{
				AdminHealth: adminHealth,
CommonHealth: commonHealth,
MasterHealth: masterHealth,
			}
		}
		
