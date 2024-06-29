//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gc-server/config/auth"
	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"

	accountHandler "github.com/game-core/gc-server/api/game/presentation/handler/account"
	exchangeHandler "github.com/game-core/gc-server/api/game/presentation/handler/exchange"
	healthHandler "github.com/game-core/gc-server/api/game/presentation/handler/health"
	loginBonusHandler "github.com/game-core/gc-server/api/game/presentation/handler/loginBonus"
	profileHandler "github.com/game-core/gc-server/api/game/presentation/handler/profile"
	authInterceptor "github.com/game-core/gc-server/api/game/presentation/interceptor/auth"
	accountUsecase "github.com/game-core/gc-server/api/game/usecase/account"
	exchangeUsecase "github.com/game-core/gc-server/api/game/usecase/exchange"
	healthUsecase "github.com/game-core/gc-server/api/game/usecase/health"
	loginBonusUsecase "github.com/game-core/gc-server/api/game/usecase/loginBonus"
	profileUsecase "github.com/game-core/gc-server/api/game/usecase/profile"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	actionService "github.com/game-core/gc-server/pkg/domain/model/action"
	eventService "github.com/game-core/gc-server/pkg/domain/model/event"
	exchangeService "github.com/game-core/gc-server/pkg/domain/model/exchange"
	googleService "github.com/game-core/gc-server/pkg/domain/model/google"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
	itemService "github.com/game-core/gc-server/pkg/domain/model/item"
	loginBonusService "github.com/game-core/gc-server/pkg/domain/model/loginBonus"
	profileService "github.com/game-core/gc-server/pkg/domain/model/profile"
	shardService "github.com/game-core/gc-server/pkg/domain/model/shard"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
	adminGoogleAuthDao "github.com/game-core/gc-server/pkg/infrastructure/auth/admin/adminGoogle"
	userItemBoxCloudWatchDao "github.com/game-core/gc-server/pkg/infrastructure/cloudwatch/user/userItemBox"
	adminHealthMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/admin/adminHealth"
	commonHealthMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonHealth"
	commonTransactionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/common/commonTransaction"
	masterActionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterAction"
	masterActionRunMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionRun"
	masterActionStepMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionStep"
	masterActionTriggerMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterActionTrigger"
	masterEventMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterEvent"
	masterExchangeMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchange"
	masterExchangeCostMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchangeCost"
	masterExchangeItemMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterExchangeItem"
	masterHealthMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterHealth"
	masterItemMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterItem"
	masterLoginBonusMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonus"
	masterLoginBonusItemMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonusItem"
	masterLoginBonusScheduleMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterLoginBonusSchedule"
	masterShardMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterShard"
	masterTransactionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/master/masterTransaction"
	userAccountMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userAccount"
	userActionMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userAction"
	userExchangeMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userExchange"
	userExchangeItemMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userExchangeItem"
	userItemBoxMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userItemBox"
	userLoginBonusMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userLoginBonus"
	userProfileMysqlDao "github.com/game-core/gc-server/pkg/infrastructure/mysql/user/userProfile"
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

func InitializeExchangeHandler() exchangeHandler.ExchangeHandler {
	wire.Build(
		exchangeHandler.NewExchangeHandler,
		InitializeExchangeUsecase,
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

func InitializeLoginBonusHandler() loginBonusHandler.LoginBonusHandler {
	wire.Build(
		loginBonusHandler.NewLoginBonusHandler,
		InitializeLoginBonusUsecase,
	)
	return nil
}

func InitializeProfileHandler() profileHandler.ProfileHandler {
	wire.Build(
		profileHandler.NewProfileHandler,
		InitializeProfileUsecase,
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

func InitializeExchangeUsecase() exchangeUsecase.ExchangeUsecase {
	wire.Build(
		exchangeUsecase.NewExchangeUsecase,
		InitializeExchangeService,
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

func InitializeLoginBonusUsecase() loginBonusUsecase.LoginBonusUsecase {
	wire.Build(
		loginBonusUsecase.NewLoginBonusUsecase,
		InitializeLoginBonusService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeProfileUsecase() profileUsecase.ProfileUsecase {
	wire.Build(
		profileUsecase.NewProfileUsecase,
		InitializeProfileService,
		InitializeTransactionService,
	)
	return nil
}

func InitializeAccountService() accountService.AccountService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		accountService.NewAccountService,
		InitializeShardService,
		userAccountMysqlDao.NewUserAccountMysqlDao,
		userAccountRedisDao.NewUserAccountRedisDao,
		userAccountTokenRedisDao.NewUserAccountTokenRedisDao,
	)
	return nil
}

func InitializeActionService() actionService.ActionService {
	wire.Build(
		database.NewMysql,
		actionService.NewActionService,
		masterActionMysqlDao.NewMasterActionMysqlDao,
		masterActionRunMysqlDao.NewMasterActionRunMysqlDao,
		masterActionStepMysqlDao.NewMasterActionStepMysqlDao,
		masterActionTriggerMysqlDao.NewMasterActionTriggerMysqlDao,
		userActionMysqlDao.NewUserActionMysqlDao,
	)
	return nil
}

func InitializeEventService() eventService.EventService {
	wire.Build(
		database.NewMysql,
		eventService.NewEventService,
		masterEventMysqlDao.NewMasterEventMysqlDao,
	)
	return nil
}

func InitializeExchangeService() exchangeService.ExchangeService {
	wire.Build(
		database.NewMysql,
		exchangeService.NewExchangeService,
		InitializeItemService,
		InitializeEventService,
		masterExchangeMysqlDao.NewMasterExchangeMysqlDao,
		masterExchangeCostMysqlDao.NewMasterExchangeCostMysqlDao,
		masterExchangeItemMysqlDao.NewMasterExchangeItemMysqlDao,
		userExchangeMysqlDao.NewUserExchangeMysqlDao,
		userExchangeItemMysqlDao.NewUserExchangeItemMysqlDao,
	)
	return nil
}

func InitializeGoogleService() googleService.GoogleService {
	wire.Build(
		auth.NewAuth,
		googleService.NewGoogleService,
		adminGoogleAuthDao.NewAdminGoogleAuthDao,
	)
	return nil
}

func InitializeHealthService() healthService.HealthService {
	wire.Build(
		database.NewMysql,
		healthService.NewHealthService,
		adminHealthMysqlDao.NewAdminHealthMysqlDao,
		commonHealthMysqlDao.NewCommonHealthMysqlDao,
		masterHealthMysqlDao.NewMasterHealthMysqlDao,
	)
	return nil
}

func InitializeItemService() itemService.ItemService {
	wire.Build(
		database.NewMysql,
		logger.NewCloudWatch,
		itemService.NewItemService,
		userItemBoxMysqlDao.NewUserItemBoxMysqlDao,
		userItemBoxCloudWatchDao.NewUserItemBoxCloudWatchDao,
		masterItemMysqlDao.NewMasterItemMysqlDao,
	)
	return nil
}

func InitializeLoginBonusService() loginBonusService.LoginBonusService {
	wire.Build(
		database.NewMysql,
		loginBonusService.NewLoginBonusService,
		InitializeItemService,
		InitializeEventService,
		userLoginBonusMysqlDao.NewUserLoginBonusMysqlDao,
		masterLoginBonusMysqlDao.NewMasterLoginBonusMysqlDao,
		masterLoginBonusItemMysqlDao.NewMasterLoginBonusItemMysqlDao,
		masterLoginBonusScheduleMysqlDao.NewMasterLoginBonusScheduleMysqlDao,
	)
	return nil
}

func InitializeProfileService() profileService.ProfileService {
	wire.Build(
		database.NewMysql,
		profileService.NewProfileService,
		userProfileMysqlDao.NewUserProfileMysqlDao,
	)
	return nil
}

func InitializeShardService() shardService.ShardService {
	wire.Build(
		database.NewMysql,
		shardService.NewShardService,
		masterShardMysqlDao.NewMasterShardMysqlDao,
	)
	return nil
}

func InitializeTransactionService() transactionService.TransactionService {
	wire.Build(
		database.NewMysql,
		database.NewRedis,
		transactionService.NewTransactionService,
		commonTransactionMysqlDao.NewCommonTransactionMysqlDao,
		masterTransactionMysqlDao.NewMasterTransactionMysqlDao,
		userTransactionMysqlDao.NewUserTransactionMysqlDao,
		userTransactionRedisDao.NewUserTransactionRedisDao,
	)
	return nil
}
