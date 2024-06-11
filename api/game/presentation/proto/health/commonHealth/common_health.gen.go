// Package commonHealth ヘルスチェック
package commonHealth

func SetCommonHealth(healthId int64, name string, commonHealthType CommonHealthType) *CommonHealth {
	return &CommonHealth{
		HealthId:         healthId,
		Name:             name,
		CommonHealthType: commonHealthType,
	}
}
