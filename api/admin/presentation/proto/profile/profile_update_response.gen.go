
// Package profile プロフィール更新レスポンス
package profile

import (
	
"github.com/game-core/gc-server/api/admin/presentation/proto/profile/userProfile"
"github.com/game-core/gc-server/api/admin/presentation/proto/profile/userProfile"
)

type ProfileUpdateResponses []*ProfileUpdateResponse

func NewProfileUpdateResponse() *ProfileUpdateResponse {
			return &ProfileUpdateResponse{}
		}

		func NewProfileUpdateResponses() ProfileUpdateResponses {
			return ProfileUpdateResponses{}
		}

		func SetProfileUpdateResponse(userProfile *userProfile.UserProfile) *ProfileUpdateResponse {
			return &ProfileUpdateResponse{
				UserProfile: userProfile,
			}
		}
		
