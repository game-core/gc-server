package google

import "context"

type GoogleRepository interface {
	GetGoogleUrl() (*GoogleURL, error)
	GetGoogleToken(ctx context.Context, code string) (*GoogleToken, error)
	GetGoogleTokenInfo(ctx context.Context, accessToken string) (*GoogleTokenInfo, error)
}
