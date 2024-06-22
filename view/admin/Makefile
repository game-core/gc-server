DOCKER_COMPOSE = docker compose -f docker-compose.local.yaml

# コンテナを起動
docker_up:
	$(DOCKER_COMPOSE) up -d --build

# コンテナに接続
view_conn:
	$(DOCKER_COMPOSE) exec view sh
