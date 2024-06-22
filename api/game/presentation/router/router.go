package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gc-server/api/game/di"
	"github.com/game-core/gc-server/api/game/presentation/proto/account"
	"github.com/game-core/gc-server/api/game/presentation/proto/exchange"
	"github.com/game-core/gc-server/api/game/presentation/proto/health"
	"github.com/game-core/gc-server/api/game/presentation/proto/loginBonus"
)

func Router(lis net.Listener) {
	// DI
	authInterceptor := di.InitializeAuthInterceptor()
	accountHandler := di.InitializeAccountHandler()
	healthHandler := di.InitializeHealthHandler()
	exchangeHandler := di.InitializeExchangeHandler()
	loginBonusHandler := di.InitializeLoginBonusHandler()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	account.RegisterAccountServer(s, accountHandler)
	health.RegisterHealthServer(s, healthHandler)
	exchange.RegisterExchangeServer(s, exchangeHandler)
	loginBonus.RegisterLoginBonusServer(s, loginBonusHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
