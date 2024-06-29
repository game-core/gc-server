// Package adminAccountGoogleUrl Google認証URL
package adminAccountGoogleUrl

type AdminAccountGoogleUrls []*AdminAccountGoogleUrl

func NewAdminAccountGoogleUrl() *AdminAccountGoogleUrl {
	return &AdminAccountGoogleUrl{}
}

func NewAdminAccountGoogleUrls() AdminAccountGoogleUrls {
	return AdminAccountGoogleUrls{}
}

func SetAdminAccountGoogleUrl(url string) *AdminAccountGoogleUrl {
	return &AdminAccountGoogleUrl{
		Url: url,
	}
}
