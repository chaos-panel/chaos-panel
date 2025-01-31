
CREATE TABLE IF NOT EXISTS `chaosplus_terminals` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `type` VARCHAR(32) NOT NULL COMMENT 'type',
    `host` VARCHAR(128) NOT NULL COMMENT 'host',
    `port` BIGINT UNSIGNED NOT NULL COMMENT 'port',
    `username` VARCHAR(128) NOT NULL COMMENT 'username',
    `password` VARCHAR(128) NOT NULL COMMENT 'password',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (`tenant_id`, `host`, `port`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'terminals';