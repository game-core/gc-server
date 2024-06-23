package router

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/game-core/gc-server/api/admin/presentation/proto/account"
	"github.com/game-core/gc-server/api/admin/presentation/proto/health"
)

func GatewayRouter(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := account.RegisterAccountHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}
	if err := health.RegisterHealthHandlerFromEndpoint(ctx, mux, endpoint, opts); err != nil {
		return err
	}

	log.Printf("gRPC gateway started")
	return nil
}
