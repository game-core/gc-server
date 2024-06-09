// Package userProfile ユーザープロフィール
package userProfile

func SetUserProfile(userId string, name string, content string) *UserProfile {
	return &UserProfile{
		UserId:  userId,
		Name:    name,
		Content: content,
	}
}
