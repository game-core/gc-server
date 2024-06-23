package account

import (
	"context"
	"errors"
	"os"

	"golang.org/x/oauth2"
	googleOAuth "golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"

	accountProto "github.com/game-core/gc-server/api/admin/presentation/proto/account"
	adminAccountGoogleTokenProto "github.com/game-core/gc-server/api/admin/presentation/proto/account/adminAccountGoogleToken"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/internal/times"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
)

type AccountUsecase interface {
	GetGoogleLoginUrl(ctx context.Context, req *accountProto.AccountGetGoogleLoginUrlRequest) (*accountProto.AccountGetGoogleLoginUrlResponse, error)
	GetGoogleLoginToken(ctx context.Context, req *accountProto.AccountGetGoogleLoginTokenRequest) (*accountProto.AccountGetGoogleLoginTokenResponse, error)
}

type accountUsecase struct {
	accountService     accountService.AccountService
	transactionService transactionService.TransactionService
}

func NewAccountUsecase(
	accountService accountService.AccountService,
	transactionService transactionService.TransactionService,
) AccountUsecase {
	return &accountUsecase{
		accountService:     accountService,
		transactionService: transactionService,
	}
}

type Google struct {
	Config *oauth2.Config
}

// GetGoogleLoginUrl アカウントをログインする
func (s *accountUsecase) GetGoogleLoginUrl(ctx context.Context, req *accountProto.AccountGetGoogleLoginUrlRequest) (*accountProto.AccountGetGoogleLoginUrlResponse, error) {
	state, err := keys.CreateStateOauthCookie()
	if err != nil {
		return nil, errors.New("接続エラー")
	}

	google := s.newGoogle()
	url := google.GetLoginURL(state)

	return accountProto.SetAccountGetGoogleLoginUrlResponse(url), nil
}

func (s *accountUsecase) GetGoogleLoginToken(ctx context.Context, req *accountProto.AccountGetGoogleLoginTokenRequest) (*accountProto.AccountGetGoogleLoginTokenResponse, error) {
	conf := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		Endpoint:     googleOAuth.Endpoint,
		Scopes:       []string{"openid", "email", "profile"},
		RedirectURL:  "http://localhost:3000",
	}

	token, err := conf.Exchange(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	return accountProto.SetAccountGetGoogleLoginTokenResponse(
		adminAccountGoogleTokenProto.SetAdminAccountGoogleToken(
			token.AccessToken,
			token.RefreshToken,
			times.TimeToPb(&token.Expiry),
		),
	), nil
}

func (s *accountUsecase) newGoogle() *Google {
	google := &Google{
		Config: &oauth2.Config{
			ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			Endpoint:     googleOAuth.Endpoint,
			Scopes:       []string{"openid", "email", "profile"},
			RedirectURL:  "http://localhost:3000",
		},
	}
	if google.Config == nil {
		panic("==== invalid key. google api ====")
	}

	return google
}

func (g *Google) GetLoginURL(state string) (clientID string) {
	return g.Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (g *Google) GetUserID(ctx context.Context, code string) (googleUserID string, err error) {
	httpClient, _ := g.Config.Exchange(ctx, code)
	if httpClient == nil {
		return "", errors.New("接続エラー")
	}

	service, err := v2.New(g.Config.Client(ctx, httpClient))
	if err != nil {
		return "", errors.New("接続エラー")
	}

	userInfo, err := service.Tokeninfo().AccessToken(httpClient.AccessToken).Context(ctx).Do()
	if err != nil {
		return "", errors.New("接続エラー")
	}

	return userInfo.UserId, nil
}
