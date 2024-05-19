CREATE TABLE master_resource
(
    master_resource_id BIGINT NOT NULL COMMENT "リソースID",
	name VARCHAR(255) NOT NULL COMMENT "リソース名",
	master_resource_enum INT NOT NULL COMMENT "リソースEnum",
	PRIMARY KEY(master_resource_id),
	UNIQUE KEY(master_resource_id),
	INDEX(master_resource_enum)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
