DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml
SCHEMA ?=

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build

# MySQLに接続
mysql_conn:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password
mysql_conn_admin:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password game_admin
mysql_conn_common:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password game_common
mysql_conn_master:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password game_master
mysql_conn_user0:
	$(DOCKER_COMPOSE) exec mysql mysql --host=localhost --user=mysql_user --password=mysql_password game_user_0
mysql_conn_user1:
	$(DOCKER_COMPOSE) exec mysql  mysql --host=localhost --user=mysql_user --password=mysql_password game_user_1

# コンテナに接続
gen_conn:
	$(DOCKER_COMPOSE) exec gen bash

# apiを生成
gen_api:
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/api/admin/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/api/game/main.go
	cd docs/proto/api/game && buf generate && cd ../../../
	$(DOCKER_COMPOSE) exec gen goimports -w .

# diを生成
gen_di:
	$(DOCKER_COMPOSE) exec gen wire api/admin/di/wire.go
	$(DOCKER_COMPOSE) exec gen wire api/game/di/wire.go

# domainを生成
gen_domain:
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/domain/model/main.go
	$(DOCKER_COMPOSE) exec gen goimports -w .

# infraを生成
gen_infra:
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/mysql/admin/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/mysql/common/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/mysql/master/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/mysql/user/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/redis/common/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/redis/user/main.go
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/pkg/infrastructure/cloudwatch/user/main.go
	$(DOCKER_COMPOSE) exec gen goimports -w .

# sqlを生成
gen_sql:
	$(DOCKER_COMPOSE) exec gen go generate ./tool/generator/sql/main.go

# mockを生成
gen_mock:
	$(DOCKER_COMPOSE) exec gen go generate ./pkg/domain/...

# マイグレーション
gen_migration:
	$(DOCKER_COMPOSE) exec gen go run ./tool/migration/migration.go

# マスターインポート
gen_master:
	$(DOCKER_COMPOSE) exec gen go run ./tool/masterImport/main.go

# フォーマット
gen_fmt:
	$(DOCKER_COMPOSE) exec gen goimports -w .
