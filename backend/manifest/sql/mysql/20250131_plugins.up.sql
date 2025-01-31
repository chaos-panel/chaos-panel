-- repos
CREATE TABLE IF NOT EXISTS `chaosplus_plugin_repos` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `name` VARCHAR(32) NOT NULL COMMENT 'name',
    `url` VARCHAR(512) NOT NULL COMMENT 'url',
    `desc` VARCHAR(512) NOT NULL COMMENT 'desc',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    INDEX (`tenant_id`),
    INDEX (`name`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'plugin repos';


-- plugins
CREATE TABLE IF NOT EXISTS `chaosplus_plugins` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `repo_id` BIGINT UNSIGNED NOT NULL COMMENT 'repo id',
    `kind` VARCHAR(32) NOT NULL COMMENT 'kind',
    `type` VARCHAR(32) NOT NULL COMMENT 'type',
    `name` VARCHAR(128) NOT NULL COMMENT 'name',
    `author` VARCHAR(128) NOT NULL COMMENT 'author',
    `version_code` BIGINT NOT NULL COMMENT 'version code',
    `version_name` VARCHAR(32) NOT NULL COMMENT 'version name',
    `desc` NVARCHAR(512) NOT NULL COMMENT 'desc',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    INDEX (`tenant_id`),
    INDEX (`kind`),
    INDEX (`type`),
    INDEX (`name`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'plugins';


-- plugin resources
CREATE TABLE IF NOT EXISTS `chaosplus_plugin_resources` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `plugin_id` BIGINT UNSIGNED NOT NULL COMMENT 'plugin id',
    `file_type` VARCHAR(16) NOT NULL COMMENT 'file type',
    `file_name` VARCHAR(128) NOT NULL COMMENT 'file name',
    `file_id` BIGINT UNSIGNED NOT NULL COMMENT 'file id',
    `mount_type` VARCHAR(16) NOT NULL COMMENT 'mount type',
    `mount_path` VARCHAR(512) NOT NULL COMMENT 'mount path',
    `mount_policy` VARCHAR(16) NOT NULL COMMENT 'mount policy',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    INDEX (`plugin_id`),
    INDEX (`file_id`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'plugin resources';

