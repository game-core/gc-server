package adminGoogle

type AdminGoogleTokenInfo struct {
	UserId        string
	Email         string
	VerifiedEmail bool
	ExpiresIn     int64
	IssuedTo      string
	Scope         string
}

func SetAdminGoogleTokenInfo(userId, email string, verifiedEmail bool, expiresIn int64, issuedTo, scope string) *AdminGoogleTokenInfo {
	return &AdminGoogleTokenInfo{
		UserId:        userId,
		Email:         email,
		VerifiedEmail: verifiedEmail,
		ExpiresIn:     expiresIn,
		IssuedTo:      issuedTo,
		Scope:         scope,
	}
}
