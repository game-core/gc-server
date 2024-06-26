CREATE TABLE master_login_bonus_schedule
(
    master_login_bonus_schedule_id BIGINT NOT NULL COMMENT "ID",
	master_login_bonus_id BIGINT NOT NULL COMMENT "ログインボーナスID",
	step INT NOT NULL COMMENT "ステップ",
	name VARCHAR(255) NOT NULL COMMENT "ログインボーナススケジュール名",
	PRIMARY KEY(master_login_bonus_schedule_id),
	UNIQUE KEY(master_login_bonus_schedule_id),
	INDEX(master_login_bonus_id),
	INDEX(step),
	INDEX(master_login_bonus_id,step)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
