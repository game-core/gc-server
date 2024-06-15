CREATE TABLE master_exchange_cost
(
    master_exchange_cost_id BIGINT NOT NULL COMMENT "交換コストID",
	master_exchange_item_id BIGINT NOT NULL COMMENT "交換アイテムID",
	master_item_id BIGINT NOT NULL COMMENT "アイテムID",
	name VARCHAR(255) NOT NULL COMMENT "交換コスト名",
	count INT NOT NULL COMMENT "個数",
	PRIMARY KEY(master_exchange_cost_id),
	UNIQUE KEY(master_exchange_cost_id),
	INDEX(master_exchange_item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
