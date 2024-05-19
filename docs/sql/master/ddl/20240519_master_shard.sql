CREATE TABLE master_shard
(
    master_shard_id BIGINT NOT NULL COMMENT "シャードID",
	shard_key VARCHAR(255) NOT NULL COMMENT "シャードキー",
	count INT NOT NULL COMMENT "シャードされたユーザー数",
	name VARCHAR(255) NOT NULL COMMENT "シャード名",
	PRIMARY KEY(master_shard_id),
	UNIQUE KEY(master_shard_id),
	INDEX(shard_key)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
