// Package commonHealth ヘルスチェック
package commonHealth

import (
	"github.com/game-core/gc-server/pkg/domain/model/health/commonHealth"
)

type CommonHealths []*CommonHealth

type CommonHealth struct {
	HealthId         int64
	Name             string
	CommonHealthEnum commonHealth.CommonHealthEnum
}

func NewCommonHealth() *CommonHealth {
	return &CommonHealth{}
}

func NewCommonHealths() CommonHealths {
	return CommonHealths{}
}

func SetCommonHealth(healthId int64, name string, commonHealthEnum commonHealth.CommonHealthEnum) *CommonHealth {
	return &CommonHealth{
		HealthId:         healthId,
		Name:             name,
		CommonHealthEnum: commonHealthEnum,
	}
}

func (t *CommonHealth) TableName() string {
	return "common_health"
}
