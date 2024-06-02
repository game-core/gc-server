// Package loginBonus ログインボーナス受け取りレスポンス
package loginBonus

import (
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus/userLoginBonus"
)

type LoginBonusReceiveResponses []*LoginBonusReceiveResponse

type LoginBonusReceiveResponse struct {
	UserLoginBonus *userLoginBonus.UserLoginBonus
}

func NewLoginBonusReceiveResponse() *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{}
}

func NewLoginBonusReceiveResponses() LoginBonusReceiveResponses {
	return LoginBonusReceiveResponses{}
}

func SetLoginBonusReceiveResponse(userLoginBonus *userLoginBonus.UserLoginBonus) *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{
		UserLoginBonus: userLoginBonus,
	}
}
