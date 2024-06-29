package adminGoogle

type AdminGoogleURL struct {
	URL string
}

func SetAdminGoogleURL(url string) *AdminGoogleURL {
	return &AdminGoogleURL{
		URL: url,
	}
}
