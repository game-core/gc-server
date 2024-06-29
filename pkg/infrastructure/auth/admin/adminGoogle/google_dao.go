package adminGoogle

import (
	"context"

	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"

	"github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/pkg/domain/model/google/adminGoogle"
)

type adminGoogleDao struct {
	GoogleConn *oauth2.Config
}

func NewAdminGoogleDao(conn *auth.AuthHandler) adminGoogle.AdminGoogleAuthRepository {
	return &adminGoogleDao{
		GoogleConn: conn.Google.Config,
	}
}

// GetAdminGoogleUrl URLを取得する
func (s *adminGoogleDao) GetAdminGoogleUrl() (*adminGoogle.AdminGoogleURL, error) {
	state, err := keys.CreateStateOauthCookie()
	if err != nil {
		return nil, errors.NewMethodError("keys.CreateStateOauthCookie", err)
	}

	return adminGoogle.SetAdminGoogleURL(s.GoogleConn.AuthCodeURL(state, oauth2.AccessTypeOffline)), nil
}

// GetAdminGoogleToken トークンを取得する
func (s *adminGoogleDao) GetAdminGoogleToken(ctx context.Context, code string) (*adminGoogle.AdminGoogleToken, error) {
	token, err := s.GoogleConn.Exchange(ctx, code)
	if err != nil {
		return nil, errors.NewMethodError("token", err)
	}

	return adminGoogle.SetAdminGoogleToken(
		token.AccessToken,
		token.RefreshToken,
		token.Expiry,
	), nil
}

// GetAdminGoogleTokenInfo トークン情報を確認する
func (s *adminGoogleDao) GetAdminGoogleTokenInfo(ctx context.Context, accessToken string) (*adminGoogle.AdminGoogleTokenInfo, error) {
	service, err := v2.New(s.GoogleConn.Client(ctx, &oauth2.Token{AccessToken: accessToken}))
	if err != nil {
		return nil, errors.NewMethodError("v2.New", err)
	}

	tokenInfo, err := service.Tokeninfo().AccessToken(accessToken).Context(ctx).Do()
	if err != nil {
		return nil, errors.NewMethodError("service.Tokeninfo", err)
	}

	return adminGoogle.SetAdminGoogleTokenInfo(
		tokenInfo.UserId,
		tokenInfo.Email,
		tokenInfo.VerifiedEmail,
		tokenInfo.ExpiresIn,
		tokenInfo.IssuedTo,
		tokenInfo.Scope,
	), nil
}
