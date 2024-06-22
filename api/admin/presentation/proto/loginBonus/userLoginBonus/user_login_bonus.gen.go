
// Package userLoginBonus ユーザーログインボーナス
package userLoginBonus

import (
	
"google.golang.org/protobuf/types/known/timestamppb"
"google.golang.org/protobuf/types/known/timestamppb"
)

type UserLoginBonuses []*UserLoginBonus

func NewUserLoginBonus() *UserLoginBonus {
			return &UserLoginBonus{}
		}

		func NewUserLoginBonuses() UserLoginBonuses {
			return UserLoginBonuses{}
		}

		func SetUserLoginBonus(userId string,masterLoginBonusId int64,receivedAt *timestamppb.Timestamp) *UserLoginBonus {
			return &UserLoginBonus{
				UserId: userId,
MasterLoginBonusId: masterLoginBonusId,
ReceivedAt: receivedAt,
			}
		}
		
