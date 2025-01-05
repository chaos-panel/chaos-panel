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
-- Table `chaosplus_tenant`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_tenant` (
    `id`           BIGINT UNSIGNED NOT NULL                                           COMMENT 'ID',
    `code`         NVARCHAR(16)    NOT NULL                                           COMMENT '代号',
    `name`         NVARCHAR(64)    NOT NULL                                           COMMENT '名称',
    `time_created` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP                 COMMENT '创建时间',
    `time_updated` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY (`code`),
    INDEX (`time_created`),
    INDEX (`time_updated`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'tenant';

-- -----------------------------------------------------
-- Table `chaosplus_worker_node`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_worker_node` (
    `id`           INT UNSIGNED AUTO_INCREMENT NOT NULL                                COMMENT '节点id',
    `port`         INT UNSIGNED                NOT NULL                                COMMENT '端口',
    `os_name`      VARCHAR(128)                NOT NULL DEFAULT ''                     COMMENT '系统名称',
    `os_arch`      VARCHAR(128)                NOT NULL DEFAULT ''                     COMMENT '系统平台',
    `os_version`   VARCHAR(128)                NOT NULL DEFAULT ''                     COMMENT '系统版本',
    `host_name`    VARCHAR(128)                NOT NULL DEFAULT ''                     COMMENT '主机名称',
    `docker`       BOOLEAN                     NOT NULL DEFAULT FALSE                  COMMENT 'IP地址',
    `ip`           VARCHAR(16)                 NOT NULL DEFAULT ''                     COMMENT 'IP地址',
    `mac`          VARCHAR(16)                 NOT NULL DEFAULT ''                     COMMENT 'MAC地址',
    `tag`          VARCHAR(32)                 NOT NULL DEFAULT ''                     COMMENT '节点标签',
    `desc`         VARCHAR(1024)               NOT NULL DEFAULT ''                     COMMENT '节点描述',
    `time_created` TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP      COMMENT '创建时间',
    `time_updated` TIMESTAMP                   NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',

    PRIMARY KEY (`id`),
    UNIQUE KEY (`tag`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'worker node';

-- -----------------------------------------------------
-- Table `chaosplus_logs`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `chaosplus_logs` (
    `id`         BIGINT UNSIGNED NOT NULL                                             COMMENT 'ID',
    `status`     VARCHAR(16)     NOT NULL DEFAULT ''                                  COMMENT '状态',
    `log`        JSON            NOT NULL                                             COMMENT '日志',
    `created_by` BIGINT UNSIGNED NOT NULL DEFAULT 0                                   COMMENT '创建者',
    `created_at` TIMESTAMP       NOT NULL DEFAULT CURRENT_TIMESTAMP                   COMMENT '创建时间',

    PRIMARY KEY (`id`),
    INDEX (`created_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci
  COMMENT = 'logs';

SHOW TABLES;