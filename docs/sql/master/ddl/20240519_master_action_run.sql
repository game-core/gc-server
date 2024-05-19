CREATE TABLE master_action_run
(
    master_action_run_id BIGINT NOT NULL COMMENT "実行されるアクションID",
	name VARCHAR(255) NOT NULL COMMENT "実行されるアクション名",
	master_action_id BIGINT NOT NULL COMMENT "アクションID",
	PRIMARY KEY(master_action_run_id),
	UNIQUE KEY(master_action_run_id),
	INDEX(master_action_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
