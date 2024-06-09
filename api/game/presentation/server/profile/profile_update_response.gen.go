// Package profile プロフィール更新レスポンス
package profile

import (
	"github.com/game-core/gc-server/api/game/presentation/server/profile/userProfile"
)

func SetProfileUpdateResponse(userProfile *userProfile.UserProfile) *ProfileUpdateResponse {
	return &ProfileUpdateResponse{
		UserProfile: userProfile,
	}
}
