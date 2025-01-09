
-- 
-- api, data, group, department, menu, button, ... tree
-- 
CREATE TABLE IF NOT EXISTS `chaosplus_resources` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `path` TEXT NOT NULL COMMENT 'path',
    `path_depth` BIGINT UNSIGNED NOT NULL COMMENT 'path depth',
    `path_md5` CHAR(32) NOT NULL COMMENT 'path md5',
    `path_sha1` CHAR(40) NOT NULL COMMENT 'path sha1',
    `path_sha256` CHAR(64) NOT NULL COMMENT 'path sha256',
    `path_sha512` CHAR(128) NOT NULL COMMENT 'path sha512',
    `children_count` BIGINT UNSIGNED NOT NULL COMMENT 'children count',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (
        `tenant_id`,
        `path_md5`,
        `path_sha1`,
        `path_sha256`,
        `path_sha512`,
        `deleted_at`
    ),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'resource';

CREATE TABLE IF NOT EXISTS `chaosplus_resources_closure` (
    `ancestor_id` BIGINT UNSIGNED NOT NULL COMMENT 'ancestor id',
    `descendant_id` BIGINT UNSIGNED NOT NULL COMMENT 'descendant id',
    `distance` BIGINT NOT NULL COMMENT 'distance',
    PRIMARY KEY (
        `ancestor_id`,
        `descendant_id`
    ),
    INDEX (`descendant_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'resource closure';


CREATE TABLE IF NOT EXISTS `chaosplus_roles` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `name` NVARCHAR (256) NOT NULL COMMENT 'name',
    `desc` TEXT NOT NULL COMMENT 'description',
    `path` TEXT NOT NULL COMMENT 'path',
    `path_depth` BIGINT UNSIGNED NOT NULL COMMENT 'path depth',
    `path_md5` CHAR(32) NOT NULL COMMENT 'path md5',
    `path_sha1` CHAR(40) NOT NULL COMMENT 'path sha1',
    `path_sha256` CHAR(64) NOT NULL COMMENT 'path sha256',
    `path_sha512` CHAR(128) NOT NULL COMMENT 'path sha512',
    `children_count` BIGINT UNSIGNED NOT NULL COMMENT 'children node count',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (
        `tenant_id`,
        `path_md5`,
        `path_sha1`,
        `path_sha256`,
        `path_sha512`,
        `deleted_at`
    ),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'role';

CREATE TABLE IF NOT EXISTS `chaosplus_roles_closure` (
    `ancestor_id` BIGINT UNSIGNED NOT NULL COMMENT 'ancestor id',
    `descendant_id` BIGINT UNSIGNED NOT NULL COMMENT 'descendant id',
    `distance` BIGINT UNSIGNED NOT NULL COMMENT 'distance',
    PRIMARY KEY (
        `ancestor_id`,
        `descendant_id`
    ),
    INDEX (`descendant_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'role closure';

CREATE TABLE IF NOT EXISTS `chaosplus_permissions` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT 'role id',
    `resource_id` BIGINT UNSIGNED NOT NULL COMMENT 'resource id',
    `action` VARCHAR(64) NOT NULL COMMENT 'action',
    `effect` VARCHAR(32) NOT NULL COMMENT 'effect', -- allow, deny
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    PRIMARY KEY (`id`),
    UNIQUE (`role_id`, `resource_id`, `action`),
    INDEX (`resource_id`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'permission';

CREATE TABLE IF NOT EXISTS `chaosplus_permissions_additional` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT 'role id',
    `type` VARCHAR(64) NOT NULL COMMENT 'type', -- IP, LOCATION
    `value` VARCHAR(64) NOT NULL COMMENT 'value', -- 192.168.1.1, 北京市
    `effect` VARCHAR(32) NOT NULL COMMENT 'effect', -- allow, deny
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    PRIMARY KEY (`id`),
    UNIQUE (`role_id`, `type`, `value`),
    INDEX (`valided_at`),
    INDEX (`expired_at`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'permission';

CREATE TABLE IF NOT EXISTS `chaosplus_organizations` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `name` VARCHAR(64) NOT NULL COMMENT 'name',
    `code` VARCHAR(64) NOT NULL COMMENT 'code',
    `logo` VARCHAR(256) NOT NULL COMMENT 'logo',
    `email` VARCHAR(128) NOT NULL COMMENT 'email',
    `location` VARCHAR(256) NOT NULL COMMENT 'location',
    `website` VARCHAR(256) NOT NULL COMMENT 'website',
    `description` VARCHAR(256) NOT NULL COMMENT 'description',
    `type` VARCHAR(32) NOT NULL COMMENT 'type', -- 个人, 企业, 政府, 学校, ...
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (`tenant_id`, `code`),
    INDEX (`tenant_id`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'organization';


CREATE TABLE IF NOT EXISTS `chaosplus_users` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `username` VARCHAR(64) NOT NULL COMMENT 'username',
    `nickname` VARCHAR(64) NOT NULL COMMENT 'nickname',
    `password` VARCHAR(128) NOT NULL COMMENT 'password',
    `phone` VARCHAR(32) NOT NULL COMMENT 'phone',
    `email` VARCHAR(128) NOT NULL COMMENT 'email',
    `wx_union_id` VARCHAR(64) NOT NULL COMMENT 'wx union id',
    `wxma_open_id` VARCHAR(64) NOT NULL COMMENT 'wxma open id',
    `wxmp_open_id` VARCHAR(64) NOT NULL COMMENT 'wxmp open id',
    `locked_by` BIGINT UNSIGNED NOT NULL COMMENT 'locked by',
    `locked_at` TIMESTAMP NOT NULL COMMENT 'locked at',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    INDEX (`tenant_id`),
    UNIQUE (`username`),
    INDEX (`locked_by`),
    INDEX (`locked_at`),
    INDEX (`created_by`),
    INDEX (`created_at`),
    INDEX (`updated_by`),
    INDEX (`updated_at`),
    INDEX (`deleted_by`),
    INDEX (`deleted_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = 'user';

CREATE TABLE IF NOT EXISTS `chaosplus_users_roles` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'user id',
    `role_id` BIGINT UNSIGNED NOT NULL COMMENT 'role id',
    `valided_at` TIMESTAMP NOT NULL COMMENT 'valided at',
    `expired_at` TIMESTAMP NOT NULL COMMENT 'expired at',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    PRIMARY KEY (`user_id`, `role_id`),
    INDEX (`valided_at`),
    INDEX (`expired_at`),
    INDEX (`created_by`),
    INDEX (`created_at`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'user roles';




CREATE TABLE IF NOT EXISTS `chaosplus_auth_clients` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `client_id` VARCHAR(32) NOT NULL COMMENT 'client id',
    `client_secret` VARCHAR(128) NOT NULL COMMENT 'client secret',
    `grant_types` TEXT NOT NULL COMMENT 'grant types',
    `scopes` TEXT NOT NULL COMMENT 'scopes',
    `redirect_uri` VARCHAR(512) NOT NULL COMMENT 'redirect uri',
    `created_by` BIGINT UNSIGNED NOT NULL COMMENT 'created by',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`client_id`),
    INDEX (`tenant_id`)
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'auth client';

CREATE TABLE IF NOT EXISTS `chaosplus_auth_tokens` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT 'ID',
    `tenant_id` BIGINT UNSIGNED NOT NULL COMMENT 'tenant id',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT 'user id',
    `client_info` TEXT NOT NULL COMMENT 'client info',
    `access_token` VARCHAR(128) NOT NULL COMMENT 'access token',
    `access_token_expired_at` TIMESTAMP NOT NULL COMMENT 'access token expires at',
    `refresh_token` VARCHAR(128) NOT NULL COMMENT 'refresh token',
    `refresh_token_expired_at` TIMESTAMP NOT NULL COMMENT 'refresh token expires at',
    `created_at` TIMESTAMP NOT NULL COMMENT 'created at',
    `updated_by` BIGINT UNSIGNED NOT NULL COMMENT 'updated by',
    `updated_at` TIMESTAMP NOT NULL COMMENT 'updated at',
    `deleted_by` BIGINT UNSIGNED NOT NULL COMMENT 'deleted by',
    `deleted_at` TIMESTAMP NOT NULL COMMENT 'deleted at',
    PRIMARY KEY (`id`),
    UNIQUE (`tenant_id`),
    INDEX (`user_id`),
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'auth token';