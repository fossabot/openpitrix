ALTER TABLE cluster_role
	ADD COLUMN mount_point VARCHAR(100) NOT NULL DEFAULT '';
ALTER TABLE cluster_role
	ADD COLUMN mount_options VARCHAR(100) NOT NULL DEFAULT '';
