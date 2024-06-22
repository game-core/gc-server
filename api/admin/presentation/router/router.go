package router

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/game-core/gc-server/api/admin/di"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health"
)

func Router(lis net.Listener) {
	// DI
	authInterceptor := di.InitializeAuthInterceptor()
	healthHandler := di.InitializeHealthHandler()

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
