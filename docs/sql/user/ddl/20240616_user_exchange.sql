CREATE TABLE user_exchange
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	master_exchange_id BIGINT NOT NULL COMMENT "交換ID",
	reset_at TIMESTAMP NOT NULL COMMENT "リセット日時",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id,master_exchange_id),
	UNIQUE KEY(user_id,master_exchange_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
