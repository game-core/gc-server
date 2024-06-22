
// Package commonHealth ヘルスチェック
package commonHealth

import (
	
)

type CommonHealths []*CommonHealth

func NewCommonHealth() *CommonHealth {
			return &CommonHealth{}
		}

		func NewCommonHealths() CommonHealths {
			return CommonHealths{}
		}

		func SetCommonHealth(healthId int64,name string,commonHealthType CommonHealthType) *CommonHealth {
			return &CommonHealth{
				HealthId: healthId,
Name: name,
CommonHealthType: commonHealthType,
			}
		}
		
