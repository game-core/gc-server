package google

type GoogleTokenInfo struct {
	UserId        string
	Email         string
	VerifiedEmail bool
	ExpiresIn     int64
	IssuedTo      string
	Scope         string
}

func SetGoogleTokenInfo(userId, email string, verifiedEmail bool, expiresIn int64, issuedTo, scope string) *GoogleTokenInfo {
	return &GoogleTokenInfo{
		UserId:        userId,
		Email:         email,
		VerifiedEmail: verifiedEmail,
		ExpiresIn:     expiresIn,
		IssuedTo:      issuedTo,
		Scope:         scope,
	}
}
