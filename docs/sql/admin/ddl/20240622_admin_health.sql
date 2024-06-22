CREATE TABLE admin_health
(
    health_id BIGINT NOT NULL COMMENT "ヘルスID",
	name VARCHAR(255) NOT NULL COMMENT "アクション名",
	admin_health_type INT NOT NULL COMMENT "ヘルスチェックタイプ",
	PRIMARY KEY(health_id),
	UNIQUE KEY(health_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
