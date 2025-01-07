-- -----------------------------------------------------
-- Init SQL
-- -----------------------------------------------------
-- SET GLOBAL tidb_skip_isolation_level_check=1;
-- SET GLOBAL tidb_multi_statement_mode='ON';
-- SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
-- DROP DATABASE IF EXISTS `chaos_plus`;
-- CREATE DATABASE `chaos_plus`;
-- USE `chaos_plus`;
-- -----------------------------------------------------
-- Table `chaosplus_worker_lock`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_distributed_lock` (
    `lock_key` VARCHAR(255) NOT NULL COMMENT 'lock key',
    `host_info` VARCHAR(2048) NOT NULL COMMENT 'host info',
    `expired_at` TIMESTAMP NOT NULL COMMENT 'expired at',
    `created_by` VARCHAR(255) NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `updated_by` VARCHAR(255) NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'time updated',
    PRIMARY KEY (`lock_key`),
    INDEX (`expired_at`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'distributed lock';
-- -----------------------------------------------------
-- Table `chaosplus_worker_id`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_worker_id` (
    `id` INT UNSIGNED NOT NULL COMMENT 'id',
    `host_info` VARCHAR(2048) NOT NULL COMMENT 'host info',
    `expired_at` TIMESTAMP NOT NULL COMMENT 'expired at',
    `created_by` VARCHAR(255) NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `updated_by` VARCHAR(255) NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'time updated',
    PRIMARY KEY (`id`),
    INDEX (`expired_at`),
    UNIQUE (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'worker id';
-- -----------------------------------------------------
-- Table `chaosplus_tenant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_tenant` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'id',
    `code` NVARCHAR (16) NOT NULL COMMENT 'code',
    `name` NVARCHAR (64) NOT NULL COMMENT 'name',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'time updated',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'time deleted',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`code`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'tenant';
-- -----------------------------------------------------
-- Table `chaosplus_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_logs` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'id',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `client_type` VARCHAR(32) NOT NULL COMMENT 'client type',
    `client_info` TEXT NOT NULL COMMENT 'client info',
    `remark` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'remark',
    `log` TEXT NOT NULL COMMENT 'log',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    PRIMARY KEY (`id`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`tenant_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'logs';
-- -----------------------------------------------------
-- Table `chaosplus_settings`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_settings` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'id',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `group` VARCHAR(32) NOT NULL COMMENT 'group',
    `key` VARCHAR(32) NOT NULL COMMENT 'key',
    `key_name` NVARCHAR (32) NOT NULL COMMENT 'key display',
    `val` TEXT NOT NULL COMMENT 'value',
    `val_type` VARCHAR(32) NOT NULL COMMENT 'value type',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'time updated',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'time deleted',
    PRIMARY KEY (`id`),
    INDEX (`created_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'settings';

-- -----------------------------------------------------
-- Table `chaosplus_snapshot`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_snapshot` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'id',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `group` VARCHAR(32) NOT NULL COMMENT 'group',
    `key` VARCHAR(32) NOT NULL COMMENT 'key',
    `key_name` NVARCHAR (32) NOT NULL COMMENT 'key display',
    `val` TEXT NOT NULL COMMENT 'value',
    `val_type` VARCHAR(32) NOT NULL COMMENT 'value type',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'time updated',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'time deleted',
    PRIMARY KEY (`id`),
    INDEX (`created_at`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'settings';