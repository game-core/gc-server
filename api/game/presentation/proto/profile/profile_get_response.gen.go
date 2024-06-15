// Package profile プロフィール作成レスポンス
package profile

import (
	"github.com/game-core/gc-server/api/game/presentation/proto/profile/userProfile"
)

func SetProfileGetResponse(userProfile *userProfile.UserProfile) *ProfileGetResponse {
	return &ProfileGetResponse{
		UserProfile: userProfile,
	}
}