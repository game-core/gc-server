//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gc-server/config/database"
	"github.com/game-core/gc-server/config/logger"

	healthHandler "github.com/game-core/gc-server/api/admin/presentation/handler/health"
	authInterceptor "github.com/game-core/gc-server/api/admin/presentation/interceptor/auth"
	healthUsecase "github.com/game-core/gc-server/api/admin/usecase/health"
	accountService "github.com/game-core/gc-server/pkg/domain/model/account"
	actionService "github.com/game-core/gc-server/pkg/domain/model/action"
	eventService "github.com/game-core/gc-server/pkg/domain/model/event"
	exchangeService "github.com/game-core/gc-server/pkg/domain/model/exchange"
	healthService "github.com/game-core/gc-server/pkg/domain/model/health"
	itemService "github.com/game-core/gc-server/pkg/domain/model/item"
	loginBonusService "github.com/game-core/gc-server/pkg/domain/model/loginBonus"
	profileService "github.com/game-core/gc-server/pkg/domain/model/profile"
	shardService "github.com/game-core/gc-server/pkg/domain/model/shard"
	transactionService "github.com/game-core/gc-server/pkg/domain/model/transaction"
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

func InitializeHealthHandler() healthHandler.HealthHandler {
	wire.Build(
		healthHandler.NewHealthHandler,
		InitializeHealthUsecase,
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

func InitializeActionService() actionService.ActionService {
	wire.Build(
		database.NewMysql,
		actionService.NewActionService,
		masterActionMysqlDao.NewMasterActionDao,
		masterActionRunMysqlDao.NewMasterActionRunDao,
		masterActionStepMysqlDao.NewMasterActionStepDao,
		masterActionTriggerMysqlDao.NewMasterActionTriggerDao,
		userActionMysqlDao.NewUserActionDao,
	)
	return nil
}

func InitializeEventService() eventService.EventService {
	wire.Build(
		database.NewMysql,
		eventService.NewEventService,
		masterEventMysqlDao.NewMasterEventDao,
	)
	return nil
}

func InitializeExchangeService() exchangeService.ExchangeService {
	wire.Build(
		database.NewMysql,
		exchangeService.NewExchangeService,
		InitializeItemService,
		InitializeEventService,
		masterExchangeMysqlDao.NewMasterExchangeDao,
		masterExchangeCostMysqlDao.NewMasterExchangeCostDao,
		masterExchangeItemMysqlDao.NewMasterExchangeItemDao,
		userExchangeMysqlDao.NewUserExchangeDao,
		userExchangeItemMysqlDao.NewUserExchangeItemDao,
	)
	return nil
}

func InitializeHealthService() healthService.HealthService {
	wire.Build(
		database.NewMysql,
		healthService.NewHealthService,
		adminHealthMysqlDao.NewAdminHealthDao,
		commonHealthMysqlDao.NewCommonHealthDao,
		masterHealthMysqlDao.NewMasterHealthDao,
	)
	return nil
}

func InitializeItemService() itemService.ItemService {
	wire.Build(
		database.NewMysql,
		logger.NewCloudWatch,
		itemService.NewItemService,
		userItemBoxMysqlDao.NewUserItemBoxDao,
		userItemBoxCloudWatchDao.NewUserItemBoxDao,
		masterItemMysqlDao.NewMasterItemDao,
	)
	return nil
}

func InitializeLoginBonusService() loginBonusService.LoginBonusService {
	wire.Build(
		database.NewMysql,
		loginBonusService.NewLoginBonusService,
		InitializeItemService,
		InitializeEventService,
		userLoginBonusMysqlDao.NewUserLoginBonusDao,
		masterLoginBonusMysqlDao.NewMasterLoginBonusDao,
		masterLoginBonusItemMysqlDao.NewMasterLoginBonusItemDao,
		masterLoginBonusScheduleMysqlDao.NewMasterLoginBonusScheduleDao,
	)
	return nil
}

func InitializeProfileService() profileService.ProfileService {
	wire.Build(
		database.NewMysql,
		profileService.NewProfileService,
		userProfileMysqlDao.NewUserProfileDao,
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
