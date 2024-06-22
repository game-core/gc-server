
// Package profile プロフィール更新リクエスト
package profile

import (
	
)

type ProfileUpdateRequests []*ProfileUpdateRequest

func NewProfileUpdateRequest() *ProfileUpdateRequest {
			return &ProfileUpdateRequest{}
		}

		func NewProfileUpdateRequests() ProfileUpdateRequests {
			return ProfileUpdateRequests{}
		}

		func SetProfileUpdateRequest(userId string,name string,content string) *ProfileUpdateRequest {
			return &ProfileUpdateRequest{
				UserId: userId,
Name: name,
Content: content,
			}
		}
		
