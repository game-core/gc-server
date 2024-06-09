// Package profile プロフィール作成レスポンス
package profile

import (
	"github.com/game-core/gc-server/api/game/presentation/server/profile/userProfile"
)

func SetProfileCreateResponse(userProfile *userProfile.UserProfile) *ProfileCreateResponse {
	return &ProfileCreateResponse{
		UserProfile: userProfile,
	}
}
