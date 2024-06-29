package auth

import (
	"context"

	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	"google.golang.org/grpc"
)

type AuthInterceptor interface {
	JwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

type authInterceptor struct {
	accountService accountService.AccountService
}

func NewAuthInterceptor(
	accountService accountService.AccountService,
) AuthInterceptor {
	return &authInterceptor{
		accountService: accountService,
	}
}

// JwtAuth 認証
func (i *authInterceptor) JwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if i.isPublic(info.FullMethod) {
		return handler(ctx, req)
	}

	return handler(ctx, req)
}

// isPublic 認証しないpath
func (i *authInterceptor) isPublic(fullMethod string) bool {
	return fullMethod == "/api.admin.Account/Login" ||
		fullMethod == "/api.admin.Account/Create" ||
		fullMethod == "/api.admin.Account/GetGoogleLoginUrl" ||
		fullMethod == "/api.admin.Account/GetGoogleLoginToken" ||
		fullMethod == "/api.admin.health/Check"
}
