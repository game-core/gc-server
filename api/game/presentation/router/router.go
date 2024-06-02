package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gc-server/api/game/di"
	"github.com/game-core/gc-server/api/game/presentation/server/account"
	"github.com/game-core/gc-server/api/game/presentation/server/health"
	"github.com/game-core/gc-server/api/game/presentation/server/loginBonus"
)

func Router(lis net.Listener) {
	// DI
	authInterceptor := di.InitializeAuthInterceptor()
	accountHandler := di.InitializeAccountHandler()
	healthHandler := di.InitializeHealthHandler()
	loginBonusHandler := di.InitializeLoginBonusHandler()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	account.RegisterAccountServer(s, accountHandler)
	health.RegisterHealthServer(s, healthHandler)
	loginBonus.RegisterLoginBonusServer(s, loginBonusHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
