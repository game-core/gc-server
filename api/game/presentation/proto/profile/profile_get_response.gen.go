
// Package profile プロフィール作成レスポンス
package profile

import (
	
"github.com/game-core/gc-server/api/game/presentation/proto/profile/userProfile"
"github.com/game-core/gc-server/api/game/presentation/proto/profile/userProfile"
)

type ProfileGetResponses []*ProfileGetResponse

func NewProfileGetResponse() *ProfileGetResponse {
			return &ProfileGetResponse{}
		}

		func NewProfileGetResponses() ProfileGetResponses {
			return ProfileGetResponses{}
		}

		func SetProfileGetResponse(userProfile *userProfile.UserProfile) *ProfileGetResponse {
			return &ProfileGetResponse{
				UserProfile: userProfile,
			}
		}
		
