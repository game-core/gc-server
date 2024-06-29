package auth

import (
	"os"

	"github.com/game-core/gc-server/internal/errors"
	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
)

var AuthHandlerInstance *AuthHandler

type AuthHandler struct {
	Google *GoogleConn
}

type GoogleConn struct {
	Config *oauth2.Config
}

// NewAuth インスタンスを作成する
func NewAuth() *AuthHandler {
	return AuthHandlerInstance
}

// InitAuth 初期化する
func InitAuth() (*AuthHandler, error) {
	authHandler := &AuthHandler{}
	if err := authHandler.auth(); err != nil {
		return nil, errors.NewMethodError("authHandler.auth", err)
	}

	AuthHandlerInstance = authHandler
	return AuthHandlerInstance, nil
}

// auth コネクションを作成する
func (s *AuthHandler) auth() error {
	google := &GoogleConn{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Endpoint:     googleOAuth.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
			RedirectURL:  os.Getenv("GC_VIEW_REDIRECT_URL"),
		},
	}
	if google.Config == nil {
		return errors.NewError("invalid key")
	}

	s.Google = google

	return nil
}
