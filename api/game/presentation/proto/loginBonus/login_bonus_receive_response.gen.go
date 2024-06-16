// Package loginBonus ログインボーナス受け取りレスポンス
package loginBonus

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/loginBonus/userLoginBonus"
)

type LoginBonusReceiveResponses []*LoginBonusReceiveResponse

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
