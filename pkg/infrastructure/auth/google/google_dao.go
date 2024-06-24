package google

import (
	"context"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"

	"github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/account/google"
)

type googleDao struct {
	GoogleConn *oauth2.Config
}

func NewGoogleDao(conn *auth.AuthHandler) google.GoogleRepository {
	return &googleDao{
		GoogleConn: conn.Google.Config,
	}
}

// GetGoogleUrl URLを取得する
func (s *googleDao) GetGoogleUrl() (*google.GoogleURL, error) {
	state, err := keys.CreateStateOauthCookie()
	if err != nil {
		return nil, errors.NewMethodError("keys.CreateStateOauthCookie", err)
	}

	return google.SetGoogleURL(s.GoogleConn.AuthCodeURL(state, oauth2.AccessTypeOffline)), nil
}

// GetGoogleToken トークンを取得する
func (s *googleDao) GetGoogleToken(ctx context.Context, code string) (*google.GoogleToken, error) {
	token, err := s.GoogleConn.Exchange(ctx, code)
	if err != nil {
		return nil, errors.NewMethodError("token", err)
	}

	return google.SetGoogleToken(
		token.AccessToken,
		token.RefreshToken,
		token.Expiry,
	), nil
}

// GetGoogleTokenInfo トークン情報を確認する
func (s *googleDao) GetGoogleTokenInfo(ctx context.Context, accessToken string) (*google.GoogleTokenInfo, error) {
	service, err := v2.New(s.GoogleConn.Client(ctx, &oauth2.Token{AccessToken: accessToken}))
	if err != nil {
		return nil, errors.NewMethodError("v2.New", err)
	}

	tokenInfo, err := service.Tokeninfo().AccessToken(accessToken).Context(ctx).Do()
	if err != nil {
		return nil, errors.NewMethodError("service.Tokeninfo", err)
	}

	return google.SetGoogleTokenInfo(
		tokenInfo.UserId,
		tokenInfo.Email,
		tokenInfo.VerifiedEmail,
		tokenInfo.ExpiresIn,
		tokenInfo.IssuedTo,
		tokenInfo.Scope,
	), nil
}
