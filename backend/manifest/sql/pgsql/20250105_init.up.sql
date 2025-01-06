-- -----------------------------------------------------
-- Init SQL
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Table `chaosplus_worker_lock`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS chaosplus_distributed_lock (
    lock_key VARCHAR(255) NOT NULL COMMENT 'lock key',
    host_info VARCHAR(2048) NOT NULL COMMENT 'host info',
    expired_at TIMESTAMP NOT NULL COMMENT 'expired at',
    created_by VARCHAR(255) NOT NULL COMMENT 'created by',
    created_at TIMESTAMP NOT NULL COMMENT 'locked at',
    updated_by VARCHAR(255) NOT NULL COMMENT 'updated by',
    updated_at TIMESTAMP NOT NULL COMMENT 'time updated',
    PRIMARY KEY (lock_key)
);
-- -----------------------------------------------------
-- Table `chaosplus_worker_id`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS chaosplus_worker_id (
    id INT UNSIGNED NOT NULL COMMENT 'id',
    host_info VARCHAR(2048) NOT NULL COMMENT 'host info',
    expired_at TIMESTAMP NOT NULL COMMENT 'expired at',
    created_by VARCHAR(255) NOT NULL COMMENT 'created by',
    created_at TIMESTAMP NOT NULL COMMENT 'locked at',
    updated_by VARCHAR(255) NOT NULL COMMENT 'updated by',
    updated_at TIMESTAMP NOT NULL COMMENT 'time updated',
    PRIMARY KEY (id),
    UNIQUE (created_by)
);
-- -----------------------------------------------------
-- Table `chaosplus_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS chaosplus_logs (
    id BIGINT UNSIGNED NOT NULL COMMENT 'id',
    err VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'err',
    log JSON NOT NULL COMMENT 'log',
    created_by BIGINT UNSIGNED NOT NULL DEFAULT 0 COMMENT 'created by',
    created_at TIMESTAMP NOT NULL COMMENT 'created at',
    PRIMARY KEY (id),
    INDEX (created_at)
);
-- -----------------------------------------------------
-- Table `chaosplus_tenant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS chaosplus_tenant (
    id BIGINT UNSIGNED NOT NULL COMMENT 'id',
    code NVARCHAR (16) NOT NULL COMMENT 'code',
    name NVARCHAR (64) NOT NULL COMMENT 'name',
    created_by BIGINT NOT NULL COMMENT 'created by',
    created_at TIMESTAMP NOT NULL COMMENT 'locked at',
    updated_by BIGINT NOT NULL COMMENT 'updated by',
    updated_at TIMESTAMP NOT NULL COMMENT 'time updated',
    PRIMARY KEY (id),
    UNIQUE (code),
    INDEX (created_at),
    INDEX (updated_at)
);

