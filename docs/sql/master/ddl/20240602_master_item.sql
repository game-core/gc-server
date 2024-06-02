CREATE TABLE master_item
(
    master_item_id BIGINT NOT NULL COMMENT "アイテムID",
	name VARCHAR(255) NOT NULL COMMENT "アイテム名",
	master_resource_enum INT NOT NULL COMMENT "リソースEnum",
	master_rarity_enum INT NOT NULL COMMENT "レアリティEnum",
	content VARCHAR(255) NOT NULL COMMENT "コンテンツ",
	PRIMARY KEY(master_item_id),
	UNIQUE KEY(master_item_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
