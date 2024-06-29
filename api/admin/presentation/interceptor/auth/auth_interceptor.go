package auth

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/game-core/gc-server/internal/errors"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
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
func (s *authInterceptor) JwtAuth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if s.isPublic(info.FullMethod) {
		return handler(ctx, req)
	}

	newHandler, err := s.setRefreshToken(ctx, req, info, handler)
	if err != nil {
		return nil, errors.NewMethodError("s.setRefreshToken", err)
	}

	return newHandler, nil
}

// isPublic 認証しないpath
func (s *authInterceptor) isPublic(fullMethod string) bool {
	return fullMethod == "/api.admin.Account/Login" ||
		fullMethod == "/api.admin.Account/Create" ||
		fullMethod == "/api.admin.Account/GetGoogleUrl" ||
		fullMethod == "/api.admin.Account/GetGoogleToken" ||
		fullMethod == "/api.admin.health/Check"
}

// setRefreshToken トークンをリフレッシュする
func (s *authInterceptor) setRefreshToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if info.FullMethod != "/api.admin.Account/RefreshGoogleToken" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.NewError("metadata is not provided")
	}

	return handler(context.WithValue(ctx, "refreshToken", strings.ReplaceAll(strings.Join(md.Get("Authorization"), " "), "Bearer ", "")), req)
}
