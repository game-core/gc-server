package google

type GoogleURL struct {
	URL string
}

func SetGoogleURL(url string) *GoogleURL {
	return &GoogleURL{
		URL: url,
	}
}
