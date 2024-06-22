
// Package profile プロフィール作成レスポンス
package profile

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/profile/userProfile"
"github.com/game-core/gc-server/api/admin/presentation/proto/profile/userProfile"
)

type ProfileCreateResponses []*ProfileCreateResponse

func NewProfileCreateResponse() *ProfileCreateResponse {
			return &ProfileCreateResponse{}
		}

		func NewProfileCreateResponses() ProfileCreateResponses {
			return ProfileCreateResponses{}
		}

		func SetProfileCreateResponse(userProfile *userProfile.UserProfile) *ProfileCreateResponse {
			return &ProfileCreateResponse{
				UserProfile: userProfile,
			}
		}
		
