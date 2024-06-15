CREATE TABLE master_exchange
(
    master_exchange_id BIGINT NOT NULL COMMENT "交換ID",
	master_event_id BIGINT NOT NULL COMMENT "イベントID",
	name VARCHAR(255) NOT NULL COMMENT "交換名",
	PRIMARY KEY(master_exchange_id),
	UNIQUE KEY(master_exchange_id),
	INDEX(master_event_id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
