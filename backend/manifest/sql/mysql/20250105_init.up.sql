-- -----------------------------------------------------
-- ChaosPanel Init SQL
-- -----------------------------------------------------

-- SET GLOBAL tidb_skip_isolation_level_check=1;
-- SET GLOBAL tidb_multi_statement_mode='ON';
-- SET GLOBAL sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

-- DROP DATABASE IF EXISTS `chaos_panel`;
-- CREATE DATABASE `chaos_panel`;
-- USE `chaos_panel`;

-- -----------------------------------------------------
-- Table `chaosplus_worker_lock`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_distributed_lock` (
    `lock_name` VARCHAR(255) PRIMARY KEY      NOT NULL                               COMMENT 'lock name',
    `locked_at` TIMESTAMP                     NOT NULL DEFAULT CURRENT_TIMESTAMP     COMMENT 'locked at',
    PRIMARY KEY (`lock_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'distributed lock';

-- -----------------------------------------------------
-- Table `chaosplus_worker_id`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_worker_id` (
    `id`           INT UNSIGNED                NOT NULL                                                       COMMENT 'id',
    `host_info`    VARCHAR(2048)               NOT NULL DEFAULT ''                                            COMMENT 'host info',
    `host_tag`     VARCHAR(128)                NOT NULL DEFAULT ''                                            COMMENT 'host tag',
    `time_created` TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP                             COMMENT 'time created',
    `time_updated` TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'time updated',

    PRIMARY KEY (`id`),
    UNIQUE KEY (`host_tag`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'worker id';

-- -----------------------------------------------------
-- Table `chaosplus_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_logs` (
    `id`         BIGINT UNSIGNED NOT NULL                                             COMMENT 'id',
    `status`     VARCHAR(16)     NOT NULL DEFAULT ''                                  COMMENT 'status',
    `log`        JSON            NOT NULL                                             COMMENT 'log',
    `created_by` BIGINT UNSIGNED NOT NULL DEFAULT 0                                   COMMENT 'created by',
    `created_at` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP                   COMMENT 'created at',

    PRIMARY KEY (`id`),
    INDEX (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'logs';


-- -----------------------------------------------------
-- Table `chaosplus_tenant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_tenant` (
    `id`           BIGINT UNSIGNED NOT NULL                                           COMMENT 'id',
    `code`         NVARCHAR(16)    NOT NULL                                           COMMENT 'code',
    `name`         NVARCHAR(64)    NOT NULL                                           COMMENT 'name',
    `time_created` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP                 COMMENT 'time created',
    `time_updated` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'time updated',

    PRIMARY KEY (`id`),
    UNIQUE KEY (`code`),
    INDEX (`time_created`),
    INDEX (`time_updated`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'tenant';


SHOW TABLES;