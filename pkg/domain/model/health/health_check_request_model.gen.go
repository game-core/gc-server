// Package health ヘルスチェックリクエスト
package health

type HealthCheckRequests []*HealthCheckRequest

type HealthCheckRequest struct {
	HealthId int64
}

func NewHealthCheckRequest() *HealthCheckRequest {
	return &HealthCheckRequest{}
}

func NewHealthCheckRequests() HealthCheckRequests {
	return HealthCheckRequests{}
}

func SetHealthCheckRequest(healthId int64) *HealthCheckRequest {
	return &HealthCheckRequest{
		HealthId: healthId,
	}
}
