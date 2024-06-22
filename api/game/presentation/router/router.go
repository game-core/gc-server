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
	s := grpc.NewServer(
		grpc.UnaryInterceptor(di.InitializeAuthInterceptor().JwtAuth),
	)

	account.RegisterAccountServer(s, di.InitializeAccountHandler())
	health.RegisterHealthServer(s, di.InitializeHealthHandler())
	exchange.RegisterExchangeServer(s, di.InitializeExchangeHandler())
	loginBonus.RegisterLoginBonusServer(s, di.InitializeLoginBonusHandler())

	serve(lis, s)
}

func serve(lis net.Listener, s *grpc.Server) {
	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
