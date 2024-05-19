CREATE TABLE master_rarity
(
    master_rarity_id BIGINT NOT NULL COMMENT "レアリティID",
	name VARCHAR(255) NOT NULL COMMENT "レアリティ名",
	master_rarity_enum INT NOT NULL COMMENT "レアリティEnum",
	PRIMARY KEY(master_rarity_id),
	UNIQUE KEY(master_rarity_id),
	INDEX(master_rarity_enum)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
