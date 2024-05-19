CREATE TABLE master_action_trigger
(
    master_action_trigger_id BIGINT NOT NULL COMMENT "アクショントリガーID",
	name VARCHAR(255) NOT NULL COMMENT "アクショントリガー名",
	master_action_trigger_enum INT NOT NULL COMMENT "アクショントリガータイプ",
	PRIMARY KEY(master_action_trigger_id),
	UNIQUE KEY(master_action_trigger_id),
	INDEX(master_action_trigger_enum)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
