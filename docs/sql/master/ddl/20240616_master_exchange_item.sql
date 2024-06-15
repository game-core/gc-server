CREATE TABLE master_exchange_item
(
    master_exchange_item_id BIGINT NOT NULL COMMENT "交換アイテムID",
	master_exchange_id BIGINT NOT NULL COMMENT "交換ID",
	master_item_id BIGINT NOT NULL COMMENT "アイテムID",
	name VARCHAR(255) NOT NULL COMMENT "交換アイテム名",
	count INT NOT NULL COMMENT "個数",
	PRIMARY KEY(master_exchange_item_id),
	UNIQUE KEY(master_exchange_item_id),
	INDEX(master_exchange_item_id),
	INDEX(master_exchange_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
