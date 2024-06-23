// Package CommonHealth ヘルスチェック
package CommonHealth

import (
	"github.com/game-core/gc-server/api/admin/presentation/proto/health/commonHealth"
)

type CommonHealths []*CommonHealth

func NewCommonHealth() *CommonHealth {
	return &CommonHealth{}
}

func NewCommonHealths() CommonHealths {
	return CommonHealths{}
}

func SetCommonHealth(healthId int64, name string, commonHealthType commonHealth.CommonHealthType) *CommonHealth {
	return &CommonHealth{
		HealthId:         healthId,
		Name:             name,
		CommonHealthType: commonHealthType,
	}
}
