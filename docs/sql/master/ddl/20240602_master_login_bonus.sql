CREATE TABLE master_login_bonus
(
    master_login_bonus_id BIGINT NOT NULL COMMENT "ログインボーナスID",
	master_event_id BIGINT NOT NULL COMMENT "イベントID",
	name VARCHAR(255) NOT NULL COMMENT "ログインボーナス名",
	PRIMARY KEY(master_login_bonus_id),
	UNIQUE KEY(master_login_bonus_id),
	INDEX(master_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
