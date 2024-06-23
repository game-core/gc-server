// Package commonHealth ヘルスチェック
package commonHealth

type CommonHealths []*CommonHealth

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
