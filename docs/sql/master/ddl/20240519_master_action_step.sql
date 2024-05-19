CREATE TABLE master_action_step
(
    master_action_step_id BIGINT NOT NULL COMMENT "アクションステップID",
	name VARCHAR(255) NOT NULL COMMENT "アクションステップ名",
	master_action_step_enum INT NOT NULL COMMENT "アクションステップEnum",
	PRIMARY KEY(master_action_step_id),
	UNIQUE KEY(master_action_step_id),
	INDEX(master_action_step_enum)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
