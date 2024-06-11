// Package loginBonus ログインボーナス受け取りレスポンス
package loginBonus

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/loginBonus/userLoginBonus"
)

func SetLoginBonusReceiveResponse(userLoginBonus *userLoginBonus.UserLoginBonus) *LoginBonusReceiveResponse {
	return &LoginBonusReceiveResponse{
		UserLoginBonus: userLoginBonus,
	}
}
