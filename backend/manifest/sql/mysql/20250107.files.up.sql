CREATE TABLE IF NOT EXISTS `chaosplus_file_storages` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `content_length` BIGINT UNSIGNED NOT NULL COMMENT 'content length',
    `content_md5` CHAR(32) NOT NULL COMMENT 'content md5',
    `content_sha1` CHAR(40) NOT NULL COMMENT 'content sha1',
    `content_sha256` CHAR(64) NOT NULL COMMENT 'content sha256',
    `content_sha512` CHAR(128) NOT NULL COMMENT 'content sha512',
    `storage_uri` TEXT NOT NULL COMMENT 'storage uri',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (
        `content_length`,
        `content_md5`,
        `content_sha1`,
        `content_sha256`,
        `content_sha512`
    ),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'file storage physical';

CREATE TABLE IF NOT EXISTS `chaosplus_file_uploading_tasks` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `expired_at` TIMESTAMP NOT NULL COMMENT 'expired at',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    PRIMARY KEY (`id`),
    INDEX (`expired_at`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'file uploading task';

CREATE TABLE IF NOT EXISTS `chaosplus_file_uploading_parts` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `expired_at` TIMESTAMP NOT NULL COMMENT 'expired at',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    PRIMARY KEY (`id`),
    INDEX (`expired_at`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'file uploading part physical';

CREATE TABLE IF NOT EXISTS `chaosplus_file_buckets` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `name` VARCHAR(128) NOT NULL COMMENT 'name',
    `capacity` BIGINT UNSIGNED NOT NULL COMMENT 'capacity',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (`name`),
    INDEX (`tenant_id`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'file logical bucket';

CREATE TABLE IF NOT EXISTS `chaosplus_files` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `bucket_id` BIGINT UNSIGNED NOT NULL COMMENT 'bucket id',
    `is_dir` TINYINT(1) UNSIGNED NOT NULL COMMENT 'is dir',
    `is_system` TINYINT(1) UNSIGNED NOT NULL DEFAULT FALSE COMMENT 'is system file',
    `is_hidden` TINYINT(1) UNSIGNED NOT NULL DEFAULT FALSE COMMENT 'is hidden file',
    `size` BIGINT UNSIGNED NOT NULL COMMENT 'size',
    `name` NVARCHAR (256) NOT NULL COMMENT 'name',
    `path` TEXT NOT NULL COMMENT 'path',
    `path_depth` BIGINT UNSIGNED NOT NULL COMMENT 'path depth',
    `path_md5` CHAR(32) NOT NULL COMMENT 'path md5',
    `path_sha1` CHAR(40) NOT NULL COMMENT 'path sha1',
    `path_sha256` CHAR(64) NOT NULL COMMENT 'path sha256',
    `path_sha512` CHAR(128) NOT NULL COMMENT 'path sha512',
    `ext` VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'extension name',
    `type` VARCHAR(32) NOT NULL DEFAULT '' COMMENT 'type',
    `subtype` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'subtype',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT 'physical file id',
    `desc` TEXT NOT NULL COMMENT 'description',
    `sub_file_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub file node count',
    `sub_file_system_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub file system node count',
    `sub_file_hidden_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub file hidden node count',
    `sub_dir_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub dir node count',
    `sub_dir_system_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub dir system node count',
    `sub_dir_hidden_count` BIGINT UNSIGNED NOT NULL COMMENT 'sub dir hidden node count',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (
        `bucket_id`,
        `path_md5`,
        `path_sha1`,
        `path_sha256`,
        `path_sha512`,
        `deleted_at`
    ),
    INDEX (`file_id`),
    INDEX (`ext`),
    INDEX (`type`, `subtype`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'file logical';

CREATE TABLE IF NOT EXISTS `chaosplus_files_closure` (
    `ancestor_id` BIGINT UNSIGNED NOT NULL COMMENT 'ancestor id',
    `descendant_id` BIGINT UNSIGNED NOT NULL COMMENT 'descendant id',
    `distance` BIGINT NOT NULL COMMENT 'distance',
    PRIMARY KEY (
        `ancestor_id`,
        `descendant_id`
    ),
    INDEX (`descendant_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'file logical closure';