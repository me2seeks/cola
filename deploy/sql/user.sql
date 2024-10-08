CREATE TABLE user (
    id bigint NOT NULL,
    create_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    delete_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    del_state tinyint NOT NULL DEFAULT 0,
    version tinyint NOT NULL DEFAULT 0,
    email char(254) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    password varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    salt char(16) NOT NULL,
    sex tinyint(1) NOT NULL DEFAULT '0' COMMENT '性别 0:男 1:女',
    avatar varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    info varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    PRIMARY KEY (id),
    UNIQUE KEY idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';


CREATE TABLE `user_auth` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `del_state` tinyint NOT NULL DEFAULT '0',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `user_id` bigint NOT NULL DEFAULT '0',
  `auth_key` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台唯一id',
  `auth_type` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '平台类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_type_key` (`auth_type`,`auth_key`) USING BTREE,
  UNIQUE KEY `idx_userId_key` (`user_id`,`auth_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户授权表';
