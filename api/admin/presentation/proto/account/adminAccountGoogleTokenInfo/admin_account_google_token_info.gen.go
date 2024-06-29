// Package adminAccountGoogleTokenInfo 管理者アカウントのGoogleToken情報
package adminAccountGoogleTokenInfo

type AdminAccountGoogleTokenInfos []*AdminAccountGoogleTokenInfo

func NewAdminAccountGoogleTokenInfo() *AdminAccountGoogleTokenInfo {
	return &AdminAccountGoogleTokenInfo{}
}

func NewAdminAccountGoogleTokenInfos() AdminAccountGoogleTokenInfos {
	return AdminAccountGoogleTokenInfos{}
}

func SetAdminAccountGoogleTokenInfo(userId string, email string, verifiedEmail bool, expiresIn int64, issuedTo string, scope string) *AdminAccountGoogleTokenInfo {
	return &AdminAccountGoogleTokenInfo{
		UserId:        userId,
		Email:         email,
		VerifiedEmail: verifiedEmail,
		ExpiresIn:     expiresIn,
		IssuedTo:      issuedTo,
		Scope:         scope,
	}
}
