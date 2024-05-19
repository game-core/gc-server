//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gc-server/config/database"

	accountHandler "github.com/game-core/gc-server/api/game/presentation/handler/account"
	healthHandler "github.com/game-core/gc-server/api/game/presentation/handler/health"
	authInterceptor "github.com/game-core/gc-server/api/game/presentation/interceptor/auth"
	accountUsecase "github.com/game-core/gc-server/api/game/usecase/account"
	healthUsecase "github.com/game-core/gc-server/api/game/usecase/health"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
	shardService "github.com/game-core/gc-server/pkg/domain/model/shard"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
	commonHealthMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonHealth"
	commonTransactionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonTransaction"
	masterHealthMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterHealth"
	masterShardMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterShard"
	masterTransactionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterTransaction"
	userAccountMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userAccount"
	userTransactionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userTransaction"
	userAccountRedisDao "github.com/game-core/gc-server/pkg/infrastructure/redis/user/userAccount"
	userAccountTokenRedisDao "github.com/game-core/gc-server/pkg/infrastructure/redis/user/userAccountToken"
	userTransactionRedisDao "github.com/game-core/gc-server/pkg/infrastructure/redis/user/userTransaction"
)

func InitializeAuthInterceptor() authInterceptor.AuthInterceptor {
	wire.Build(
		authInterceptor.NewAuthInterceptor,
		InitializeAccountService,
	)
	return nil
}

func InitializeAccountHandler() accountHandler.AccountHandler {
	wire.Build(
		accountHandler.NewAccountHandler,
		InitializeAccountUsecase,
	)
	return nil
}

func InitializeHealthHandler() healthHandler.HealthHandler {
	wire.Build(
		healthHandler.NewHealthHandler,
		InitializeHealthUsecase,
	)
	return nil
}

func InitializeAccountUsecase() accountUsecase.AccountUsecase {
	wire.Build(
		accountUsecase.NewAccountUsecase,
		InitializeAccountService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeHealthUsecase() healthUsecase.HealthUsecase {
	wire.Build(
		healthUsecase.NewHealthUsecase,
		InitializeHealthService,
	)
	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		accountService.NewAccountService,
		InitializeShardService,
		userAccountMysqlDao.NewUserAccountDao,
		userAccountRedisDao.NewUserAccountDao,
		userAccountTokenRedisDao.NewUserAccountTokenDao,
	)
	return nil
}

func InitializeHealthService() healthService.HealthService {
	wire.Build(
		database.NewMysql,
		healthService.NewHealthService,
		commonHealthMysqlDao.NewCommonHealthDao,
		masterHealthMysqlDao.NewMasterHealthDao,
	)
	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		database.NewMysql,
		shardService.NewShardService,
		masterShardMysqlDao.NewMasterShardDao,
	)
	return nil
}

func InitializeTransactionService() transactionService.TransactionService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		transactionService.NewTransactionService,
		commonTransactionMysqlDao.NewCommonTransactionDao,
		masterTransactionMysqlDao.NewMasterTransactionDao,
		userTransactionMysqlDao.NewUserTransactionDao,
		userTransactionRedisDao.NewUserTransactionDao,
	)
	return nil
}
