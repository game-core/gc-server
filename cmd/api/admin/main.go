package main

import (
	"context"
	"flag"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/game-core/gc-server/api/admin/presentation/router"
	"github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"
)

func main() {
	if _, err := auth.InitAuth(); err != nil {
		log.Fatalf("failed to auth.InitAuth: %v", err)
	}

	if _, err := database.InitMysql(); err != nil {
		log.Fatalf("failed to database.InitMysql: %v", err)
	}

	if _, err := database.InitRedis(); err != nil {
		log.Fatalf("failed to database.InitRedis: %v", err)
	}

	if _, err := logger.InitCloudWatch(); err != nil {
		log.Fatalf("failed to logger.InitCloudWatch: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to net.Listen: %v", err)
	}
	defer func(lis net.Listener) {
		if err := lis.Close(); err != nil {
			log.Fatalf("failed to lis.Close: %v", err)
		}
	}(lis)

	runGroup(lis)
}

func runGroup(lis net.Listener) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		router.Router(lis)
	}()
	go func() {
		defer wg.Done()
		gateway()
	}()
	wg.Wait()
}

func gateway() {
	flag.Parse()
	defer glog.Flush()
	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := router.GatewayRouter(ctx, mux, *flag.String("grpc-server-endpoint", "localhost:50051", "gRPC server endpoint"), opts); err != nil {
		return err
	}

	return http.ListenAndServe(":8001", mux)
}
