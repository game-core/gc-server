// Package userProfile ユーザープロフィール
package userProfile

type UserProfiles []*UserProfile

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}

func NewUserProfiles() UserProfiles {
	return UserProfiles{}
}

func SetUserProfile(userId string, name string, content string) *UserProfile {
	return &UserProfile{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
