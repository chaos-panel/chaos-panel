-- -----------------------------------------------------
-- Init SQL
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Table `chaosplus_worker_lock`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_distributed_lock` (
    `lock_key` VARCHAR(255) NOT NULL PRIMARY KEY COMMENT 'lock key',
    `host_info` VARCHAR(2048) NOT NULL COMMENT 'host info',
    `expired_at` DATETIME NOT NULL COMMENT 'expired at',
    `created_by` VARCHAR(255) NOT NULL COMMENT 'created by',
    `created_at` DATETIME NOT NULL COMMENT 'locked at',
    `updated_by` VARCHAR(255) NOT NULL COMMENT 'updated by',
    `updated_at` DATETIME NOT NULL COMMENT 'time updated'
);
-- -----------------------------------------------------
-- Table `chaosplus_worker_id`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_worker_id` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT COMMENT 'id',
    `host_info` VARCHAR(2048) NOT NULL COMMENT 'host info',
    `expired_at` DATETIME NOT NULL COMMENT 'expired at',
    `created_by` VARCHAR(255) NOT NULL COMMENT 'created by',
    `created_at` DATETIME NOT NULL COMMENT 'locked at',
    `updated_by` VARCHAR(255) NOT NULL COMMENT 'updated by',
    `updated_at` DATETIME NOT NULL COMMENT 'time updated',
    UNIQUE (`created_by`)
);
-- -----------------------------------------------------
-- Table `chaosplus_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_logs` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT COMMENT 'id',
    `err` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'err',
    `log` TEXT NOT NULL COMMENT 'log',
    `created_by` INTEGER NOT NULL DEFAULT 0 COMMENT 'created by',
    `created_at` DATETIME NOT NULL COMMENT 'created at'
);
-- -----------------------------------------------------
-- Table `chaosplus_tenant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_tenant` (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT COMMENT 'id',
    `code` NVARCHAR (16) NOT NULL COMMENT 'code',
    `name` NVARCHAR (64) NOT NULL COMMENT 'name',
    `created_by` INTEGER NOT NULL COMMENT 'created by',
    `created_at` DATETIME NOT NULL COMMENT 'locked at',
    `updated_by` INTEGER NOT NULL COMMENT 'updated by',
    `updated_at` DATETIME NOT NULL COMMENT 'time updated',
    UNIQUE (`code`)
);