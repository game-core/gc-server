//go:generate mockgen -source=./account_service.go -destination=./account_service_mock.gen.go -package=account
package account

import (
	"context"

	"github.com/game-core/gc-server/internal/errors"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccount"
	"github.com/game-core/gc-server/pkg/domain/model/account/userAccountToken"
)

type AccountService interface {
	CheckToken(ctx context.Context, req *AccountCheckTokenRequest) (*AccountCheckTokenResponse, error)
}

type accountService struct {
	userAccountMysqlRepository      userAccount.UserAccountMysqlRepository
	userAccountTokenRedisRepository userAccountToken.UserAccountTokenRedisRepository
}

func NewAccountService(
	userAccountMysqlRepository userAccount.UserAccountMysqlRepository,
	userAccountTokenRedisRepository userAccountToken.UserAccountTokenRedisRepository,
) AccountService {
	return &accountService{
		userAccountMysqlRepository:      userAccountMysqlRepository,
		userAccountTokenRedisRepository: userAccountTokenRedisRepository,
	}
}

// CheckToken トークンを検証する
func (s *accountService) CheckToken(ctx context.Context, req *AccountCheckTokenRequest) (*AccountCheckTokenResponse, error) {
	userAccountTokenModel, err := s.userAccountTokenRedisRepository.Find(ctx, req.UserId)
	if err != nil {
		return nil, errors.NewMethodError("s.userAccountTokenRedisRepository.Find", err)
	}

	return SetAccountCheckTokenResponse(userAccountTokenModel), nil
}
