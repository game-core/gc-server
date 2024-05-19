package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gc-server/api/game/di"
	"github.com/game-core/gc-server/api/game/presentation/server/health"
)

func Router(lis net.Listener) {
	// DI
	healthHandler := di.InitializeHealthHandler()
	authInterceptor := di.InitializeAuthInterceptor()

	// Server
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.JwtAuth),
	)

	health.RegisterHealthServer(s, healthHandler)

	log.Printf("gRPC server started")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
