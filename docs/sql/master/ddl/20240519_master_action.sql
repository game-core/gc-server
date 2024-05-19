CREATE TABLE master_action
(
    master_action_id BIGINT NOT NULL COMMENT "アクションID",
	name VARCHAR(255) NOT NULL COMMENT "アクション名",
	master_action_step_enum INT NOT NULL COMMENT "アクションステップタイプ",
	master_action_trigger_enum INT NOT NULL COMMENT "アクショントリガータイプ",
	target_id BIGINT DEFAULT NULL COMMENT "対象のID",
	trigger_master_action_id BIGINT DEFAULT NULL COMMENT "トリガーになるアクションのID",
	expiration INT DEFAULT NULL COMMENT "有効期限",
	PRIMARY KEY(master_action_id),
	UNIQUE KEY(master_action_id),
	INDEX(master_action_step_enum),
	INDEX(target_id),
	INDEX(master_action_step_enum,target_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
