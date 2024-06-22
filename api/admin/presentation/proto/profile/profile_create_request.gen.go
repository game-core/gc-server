
// Package profile プロフィール作成リクエスト
package profile

import (
	
)

type ProfileCreateRequests []*ProfileCreateRequest

func NewProfileCreateRequest() *ProfileCreateRequest {
			return &ProfileCreateRequest{}
		}

		func NewProfileCreateRequests() ProfileCreateRequests {
			return ProfileCreateRequests{}
		}

		func SetProfileCreateRequest(userId string,name string,content string) *ProfileCreateRequest {
			return &ProfileCreateRequest{
				UserId: userId,
Name: name,
Content: content,
			}
		}
		
