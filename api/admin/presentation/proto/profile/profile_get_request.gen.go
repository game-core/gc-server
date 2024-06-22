
// Package profile プロフィール取得リクエスト
package profile

import (
	
)

type ProfileGetRequests []*ProfileGetRequest

func NewProfileGetRequest() *ProfileGetRequest {
			return &ProfileGetRequest{}
		}

		func NewProfileGetRequests() ProfileGetRequests {
			return ProfileGetRequests{}
		}

		func SetProfileGetRequest(userId string) *ProfileGetRequest {
			return &ProfileGetRequest{
				UserId: userId,
			}
		}
		
