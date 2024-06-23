// Package commonHealth ヘルスチェック
package commonHealth

type CommonHealths []*CommonHealth

type CommonHealth struct {
	HealthId         int64
	Name             string
	CommonHealthEnum CommonHealthEnum
}

func NewCommonHealth() *CommonHealth {
	return &CommonHealth{}
}

func NewCommonHealths() CommonHealths {
	return CommonHealths{}
}

func SetCommonHealth(healthId int64, name string, commonHealthEnum CommonHealthEnum) *CommonHealth {
	return &CommonHealth{
		HealthId:         healthId,
		Name:             name,
		CommonHealthEnum: commonHealthEnum,
	}
}
