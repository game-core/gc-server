CREATE TABLE user_exchange_item
(
    user_id VARCHAR(255) NOT NULL COMMENT "ユーザーID",
	master_exchange_id BIGINT NOT NULL COMMENT "交換ID",
	master_exchange_item_id BIGINT NOT NULL COMMENT "交換アイテムID",
	count INT NOT NULL COMMENT "残りのアイテム数",
	created_at TIMESTAMP NOT NULL COMMENT "作成日時",
	updated_at TIMESTAMP NOT NULL COMMENT "更新日時",
	PRIMARY KEY(user_id,master_exchange_item_id),
	UNIQUE KEY(user_id,master_exchange_item_id),
	INDEX(user_id,master_exchange_id),
	INDEX(user_id,master_exchange_item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
