package main

import (
	"github.com/game-core/gc-server/api/admin/presentation/router"
	"github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"
	"log"
	"net"
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

	router.Router(lis)
}
