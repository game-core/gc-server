package main

import (
	"fmt"
	"log"
	"net"

	"github.com/game-core/gc-server/api/admin/presentation/router"
	apiConfig "github.com/game-core/gc-server/config/api"
	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"
)

func main() {
	if _, err := database.InitMysql(); err != nil {
		log.Fatalf("failed to database.InitMysql: %v", err)
	}

	if _, err := database.InitRedis(); err != nil {
		log.Fatalf("failed to database.InitRedis: %v", err)
	}

	if _, err := logger.InitCloudWatch(); err != nil {
		log.Fatalf("failed to logger.InitCloudWatch: %v", err)
	}

	apiConfig := apiConfig.GetAppConfig()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", apiConfig.Port.GrpcPort))
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
