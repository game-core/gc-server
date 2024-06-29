// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/game-core/gc-server/api/admin/presentation/handler/account"
	"github.com/game-core/gc-server/api/admin/presentation/handler/health"
	"github.com/game-core/gc-server/api/admin/presentation/interceptor/auth"
	account2 "github.com/game-core/gc-server/api/admin/usecase/account"
	health2 "github.com/game-core/gc-server/api/admin/usecase/health"
	auth2 "github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"
	account3 "github.com/game-core/gc-server/pkg/domain/model/account"
	"github.com/game-core/gc-server/pkg/domain/model/action"
	"github.com/game-core/gc-server/pkg/domain/model/event"
	"github.com/game-core/gc-server/pkg/domain/model/exchange"
	"github.com/game-core/gc-server/pkg/domain/model/google"
	health3 "github.com/game-core/gc-server/pkg/domain/model/health"
	"github.com/game-core/gc-server/pkg/domain/model/item"
	"github.com/game-core/gc-server/pkg/domain/model/loginBonus"
	"github.com/game-core/gc-server/pkg/domain/model/profile"
	"github.com/game-core/gc-server/pkg/domain/model/shard"
	"github.com/game-core/gc-server/pkg/domain/model/transaction"
	"github.com/game-core/gc-server/pkg/infrastructure/auth/admin/adminGoogle"
	userItemBox2 "github.com/game-core/gc-server/pkg/infrastructure/cloudwatch/user/userItemBox"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/admin/adminHealth"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonHealth"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonTransaction"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterAction"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionRun"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionStep"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionTrigger"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterEvent"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchange"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchangeCost"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchangeItem"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterHealth"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterItem"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonus"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterShard"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterTransaction"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userAccount"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userAction"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userExchange"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userExchangeItem"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userItemBox"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userLoginBonus"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userProfile"
	"github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userTransaction"
	userAccount2 "github.com/game-core/gc-server/pkg/infrastructure/redis/user/userAccount"
	"github.com/game-core/gc-server/pkg/infrastructure/redis/user/userAccountToken"
	masterTransaction2 "github.com/game-core/gc-server/pkg/infrastructure/redis/user/userTransaction"
)

// Injectors from wire.go:

func InitializeAuthInterceptor() auth.AuthInterceptor {
	accountService := InitializeAccountService()
	authInterceptor := auth.NewAuthInterceptor(accountService)
	return authInterceptor
}

func InitializeAccountHandler() account.AccountHandler {
	accountUsecase := InitializeAccountUsecase()
	accountHandler := account.NewAccountHandler(accountUsecase)
	return accountHandler
}

func InitializeHealthHandler() health.HealthHandler {
	healthUsecase := InitializeHealthUsecase()
	healthHandler := health.NewHealthHandler(healthUsecase)
	return healthHandler
}

func InitializeAccountUsecase() account2.AccountUsecase {
	accountService := InitializeAccountService()
	accountUsecase := account2.NewAccountUsecase(accountService)
	return accountUsecase
}

func InitializeHealthUsecase() health2.HealthUsecase {
	healthService := InitializeHealthService()
	healthUsecase := health2.NewHealthUsecase(healthService)
	return healthUsecase
}

func InitializeAccountService() account3.AccountService {
	shardService := InitializeShardService()
	googleService := InitializeGoogleService()
	mysqlHandler := database.NewMysql()
	userAccountMysqlRepository := userAccount.NewUserAccountMysqlDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userAccountRedisRepository := userAccount2.NewUserAccountRedisDao(redisHandler)
	userAccountTokenRedisRepository := userAccountToken.NewUserAccountTokenRedisDao(redisHandler)
	accountService := account3.NewAccountService(shardService, googleService, userAccountMysqlRepository, userAccountRedisRepository, userAccountTokenRedisRepository)
	return accountService
}

func InitializeActionService() action.ActionService {
	mysqlHandler := database.NewMysql()
	masterActionMysqlRepository := masterAction.NewMasterActionMysqlDao(mysqlHandler)
	masterActionRunMysqlRepository := masterActionRun.NewMasterActionRunMysqlDao(mysqlHandler)
	masterActionStepMysqlRepository := masterActionStep.NewMasterActionStepMysqlDao(mysqlHandler)
	masterActionTriggerMysqlRepository := masterActionTrigger.NewMasterActionTriggerMysqlDao(mysqlHandler)
	userActionMysqlRepository := userAction.NewUserActionMysqlDao(mysqlHandler)
	actionService := action.NewActionService(masterActionMysqlRepository, masterActionRunMysqlRepository, masterActionStepMysqlRepository, masterActionTriggerMysqlRepository, userActionMysqlRepository)
	return actionService
}

func InitializeEventService() event.EventService {
	mysqlHandler := database.NewMysql()
	masterEventMysqlRepository := masterEvent.NewMasterEventMysqlDao(mysqlHandler)
	eventService := event.NewEventService(masterEventMysqlRepository)
	return eventService
}

func InitializeExchangeService() exchange.ExchangeService {
	itemService := InitializeItemService()
	eventService := InitializeEventService()
	mysqlHandler := database.NewMysql()
	masterExchangeMysqlRepository := masterExchange.NewMasterExchangeMysqlDao(mysqlHandler)
	masterExchangeCostMysqlRepository := masterExchangeCost.NewMasterExchangeCostMysqlDao(mysqlHandler)
	masterExchangeItemMysqlRepository := masterExchangeItem.NewMasterExchangeItemMysqlDao(mysqlHandler)
	userExchangeMysqlRepository := userExchange.NewUserExchangeMysqlDao(mysqlHandler)
	userExchangeItemMysqlRepository := userExchangeItem.NewUserExchangeItemMysqlDao(mysqlHandler)
	exchangeService := exchange.NewExchangeService(itemService, eventService, masterExchangeMysqlRepository, masterExchangeCostMysqlRepository, masterExchangeItemMysqlRepository, userExchangeMysqlRepository, userExchangeItemMysqlRepository)
	return exchangeService
}

func InitializeGoogleService() google.GoogleService {
	authHandler := auth2.NewAuth()
	adminGoogleAuthRepository := adminGoogle.NewAdminGoogleAuthDao(authHandler)
	googleService := google.NewGoogleService(adminGoogleAuthRepository)
	return googleService
}

func InitializeHealthService() health3.HealthService {
	mysqlHandler := database.NewMysql()
	adminHealthMysqlRepository := adminHealth.NewAdminHealthMysqlDao(mysqlHandler)
	commonHealthMysqlRepository := commonHealth.NewCommonHealthMysqlDao(mysqlHandler)
	masterHealthMysqlRepository := masterHealth.NewMasterHealthMysqlDao(mysqlHandler)
	healthService := health3.NewHealthService(adminHealthMysqlRepository, commonHealthMysqlRepository, masterHealthMysqlRepository)
	return healthService
}

func InitializeItemService() item.ItemService {
	mysqlHandler := database.NewMysql()
	userItemBoxMysqlRepository := userItemBox.NewUserItemBoxMysqlDao(mysqlHandler)
	cloudWatchHandler := logger.NewCloudWatch()
	userItemBoxCloudWatchRepository := userItemBox2.NewUserItemBoxCloudWatchDao(cloudWatchHandler)
	masterItemMysqlRepository := masterItem.NewMasterItemMysqlDao(mysqlHandler)
	itemService := item.NewItemService(userItemBoxMysqlRepository, userItemBoxCloudWatchRepository, masterItemMysqlRepository)
	return itemService
}

func InitializeLoginBonusService() loginBonus.LoginBonusService {
	itemService := InitializeItemService()
	eventService := InitializeEventService()
	mysqlHandler := database.NewMysql()
	userLoginBonusMysqlRepository := userLoginBonus.NewUserLoginBonusMysqlDao(mysqlHandler)
	masterLoginBonusMysqlRepository := masterLoginBonus.NewMasterLoginBonusMysqlDao(mysqlHandler)
	masterLoginBonusItemMysqlRepository := masterLoginBonusItem.NewMasterLoginBonusItemMysqlDao(mysqlHandler)
	masterLoginBonusScheduleMysqlRepository := masterLoginBonusSchedule.NewMasterLoginBonusScheduleMysqlDao(mysqlHandler)
	loginBonusService := loginBonus.NewLoginBonusService(itemService, eventService, userLoginBonusMysqlRepository, masterLoginBonusMysqlRepository, masterLoginBonusItemMysqlRepository, masterLoginBonusScheduleMysqlRepository)
	return loginBonusService
}

func InitializeProfileService() profile.ProfileService {
	mysqlHandler := database.NewMysql()
	userProfileMysqlRepository := userProfile.NewUserProfileMysqlDao(mysqlHandler)
	profileService := profile.NewProfileService(userProfileMysqlRepository)
	return profileService
}

func InitializeShardService() shard.ShardService {
	mysqlHandler := database.NewMysql()
	masterShardMysqlRepository := masterShard.NewMasterShardMysqlDao(mysqlHandler)
	shardService := shard.NewShardService(masterShardMysqlRepository)
	return shardService
}

func InitializeTransactionService() transaction.TransactionService {
	mysqlHandler := database.NewMysql()
	commonTransactionMysqlRepository := commonTransaction.NewCommonTransactionMysqlDao(mysqlHandler)
	masterTransactionMysqlRepository := masterTransaction.NewMasterTransactionMysqlDao(mysqlHandler)
	userTransactionMysqlRepository := userTransaction.NewUserTransactionMysqlDao(mysqlHandler)
	redisHandler := database.NewRedis()
	userTransactionRedisRepository := masterTransaction2.NewUserTransactionRedisDao(redisHandler)
	transactionService := transaction.NewTransactionService(commonTransactionMysqlRepository, masterTransactionMysqlRepository, userTransactionMysqlRepository, userTransactionRedisRepository)
	return transactionService
}
