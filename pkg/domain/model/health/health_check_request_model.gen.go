// Package health ヘルスチェックリクエスト
package health

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
	"github.com/game-core/gc-server/pkg/domain/model/health/masterHealth"
)

type HealthCheckRequests []*HealthCheckRequest

type HealthCheckRequest struct {
	HealthId         int64
	Name             string
	CommonHealthType commonHealth.CommonHealthType
	MasterHealthType masterHealth.MasterHealthType
}

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
