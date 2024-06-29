//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/internal/keys"
	"github.com/game-core/gc-server/internal/times"
	"github.com/game-core/gc-server/internal/tokens"
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleToken"
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleTokenInfo"
	"github.com/game-core/gc-server/pkg/domain/model/account/adminAccountGoogleUrl"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccountToken"
	"github.com/game-core/gc-server/pkg/domain/model/google"
	"github.com/game-core/gc-server/pkg/domain/model/shard"
)

type AccountService interface {
	Get(ctx context.Context, req *AccountGetRequest) (*AccountGetResponse, error)
	GetToken(ctx context.Context, req *AccountGetTokenRequest) (*AccountGetTokenResponse, error)
	Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error)
	CreateUserId(ctx context.Context) (string, error)
	Login(ctx context.Context, mtx *gorm.DB, rtx redis.Pipeliner, req *AccountLoginRequest) (*AccountLoginResponse, error)
	GetGoogleUrl() (*AccountGetGoogleUrlResponse, error)
	GetGoogleToken(ctx context.Context, req *AccountGetGoogleTokenRequest) (*AccountGetGoogleTokenResponse, error)
	GetGoogleTokenInfo(ctx context.Context, req *AccountGetGoogleTokenInfoRequest) (*AccountGetGoogleTokenInfoResponse, error)
}

type accountService struct {
	shardService                    shard.ShardService
	googleService                   google.GoogleService
	userAccountMysqlRepository      userAccount.UserAccountMysqlRepository
	userAccountRedisRepository      userAccount.UserAccountRedisRepository
	userAccountTokenRedisRepository userAccountToken.UserAccountTokenRedisRepository
}

func NewAccountService(
	shardService shard.ShardService,
	googleService google.GoogleService,
	userAccountMysqlRepository userAccount.UserAccountMysqlRepository,
	userAccountRedisRepository userAccount.UserAccountRedisRepository,
	userAccountTokenRedisRepository userAccountToken.UserAccountTokenRedisRepository,
) AccountService {
	return &accountService{
		shardService:                    shardService,
		googleService:                   googleService,
		userAccountMysqlRepository:      userAccountMysqlRepository,
		userAccountRedisRepository:      userAccountRedisRepository,
		userAccountTokenRedisRepository: userAccountTokenRedisRepository,
	}
}

// Get ユーザーを確認する
func (s *accountService) Get(ctx context.Context, req *AccountGetRequest) (*AccountGetResponse, error) {
	userAccountModel, err := s.userAccountRedisRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountRedisRepository.Find", err)
	}

	return SetAccountGetResponse(userAccountModel), err
}

// GetToken トークンを検証する
func (s *accountService) GetToken(ctx context.Context, req *AccountGetTokenRequest) (*AccountGetTokenResponse, error) {
	userAccountTokenModel, err := s.userAccountTokenRedisRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountTokenRedisRepository.Find", err)
	}

	return SetAccountGetTokenResponse(userAccountTokenModel), nil
}

// Create アカウントを作成する
func (s *accountService) Create(ctx context.Context, tx *gorm.DB, req *AccountCreateRequest) (*AccountCreateResponse, error) {
	hashPassword, err := keys.CreateHashPassword(req.Password)
	if err != nil {
		return nil, errors.NewMethodError("keys.CreateHashPassword", err)
	}

	userAccountModel, err := s.userAccountMysqlRepository.Create(ctx, tx, userAccount.SetUserAccount(req.UserId, req.Name, hashPassword, times.Now(), times.Now()))
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Create", err)
	}

	return SetAccountCreateResponse(userAccountModel), nil
}

// CreateUserId ユーザーIDを作成する
func (s *accountService) CreateUserId(ctx context.Context) (string, error) {
	shardKey, err := s.shardService.GetShardKey(ctx)
	if err != nil {
		return "", errors.NewMethodError("s.shardService.GetShardKey", err)
	}

	userId, err := keys.CreateUserId(shardKey)
	if err != nil {
		return "", errors.NewMethodError("keys.GenerateUserId", err)
	}

	return userId, nil
}

// Login ログインする
func (s *accountService) Login(ctx context.Context, mtx *gorm.DB, rtx redis.Pipeliner, req *AccountLoginRequest) (*AccountLoginResponse, error) {
	userAccountModel, err := s.userAccountMysqlRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Find", err)
	}

	if !keys.CheckPassword(req.Password, userAccountModel.Password) {
		return nil, errors.NewError("invalid password")
	}

	userAccountModel.LoginAt = times.Now()
	result, err := s.userAccountMysqlRepository.Update(ctx, mtx, userAccountModel)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountMysqlRepository.Update", err)
	}

	if _, err := s.userAccountRedisRepository.Set(ctx, rtx, userAccountModel); err != nil {
		return nil, errors.NewMethodError("s.userAccountRedisRepository.Set", err)
	}

	token, err := tokens.GenerateAuthTokenByUserId(userAccountModel.UserId, userAccountModel.Name)
	if err != nil {
		return nil, errors.NewMethodError("tokens.GenerateAuthTokenByUserId", err)
	}

	if _, err := s.userAccountTokenRedisRepository.Set(ctx, rtx, userAccountToken.SetUserAccountToken(userAccountModel.UserId, token)); err != nil {
		return nil, errors.NewMethodError("s.userAccountTokenRedisRepository.Set", err)
	}

	return SetAccountLoginResponse(token, result), nil
}

// GetGoogleUrl 認証URLを取得する
func (s *accountService) GetGoogleUrl() (*AccountGetGoogleUrlResponse, error) {
	adminGoogleModel, err := s.googleService.GetAdminGoogleUrl()
	if err != nil {
		return nil, errors.NewMethodError("s.googleService.GetAdminGoogleUrl", err)
	}

	return SetAccountGetGoogleUrlResponse(
		adminAccountGoogleUrl.SetAdminAccountGoogleUrl(
			adminGoogleModel.URL,
		),
	), nil
}

// GetGoogleToken トークンを取得する
func (s *accountService) GetGoogleToken(ctx context.Context, req *AccountGetGoogleTokenRequest) (*AccountGetGoogleTokenResponse, error) {
	adminGoogleModel, err := s.googleService.GetAdminGoogleToken(ctx, req.Code)
	if err != nil {
		return nil, errors.NewMethodError("s.googleService.GetAdminGoogleToken", err)
	}

	return SetAccountGetGoogleTokenResponse(
		adminAccountGoogleToken.SetAdminAccountGoogleToken(
			adminGoogleModel.AccessToken,
			adminGoogleModel.RefreshToken,
			adminGoogleModel.ExpiredAt,
		),
	), nil
}

// GetGoogleTokenInfo トークン情報を取得する
func (s *accountService) GetGoogleTokenInfo(ctx context.Context, req *AccountGetGoogleTokenInfoRequest) (*AccountGetGoogleTokenInfoResponse, error) {
	adminGoogleModel, err := s.googleService.GetAdminGoogleTokenInfo(ctx, req.AccessToken)
	if err != nil {
		return nil, errors.NewMethodError("s.googleService.GetAdminGoogleTokenInfo", err)
	}

	return SetAccountGetGoogleTokenInfoResponse(
		adminAccountGoogleTokenInfo.SetAdminAccountGoogleTokenInfo(
			adminGoogleModel.UserId,
			adminGoogleModel.Email,
			adminGoogleModel.VerifiedEmail,
			adminGoogleModel.ExpiresIn,
			adminGoogleModel.IssuedTo,
			adminGoogleModel.Scope,
		),
	), nil
}
