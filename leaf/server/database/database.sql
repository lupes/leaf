CREATE TABLE `problem`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `title`        varchar(255)        NOT NULL COMMENT '题目标题',
    `content`      text                NOT NULL COMMENT '题目内容',
    `created_time` datetime            NOT NULL COMMENT '创建时间',
    `updated_time` datetime            NOT NULL COMMENT '更新时间',
    `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 16
  DEFAULT CHARSET = utf8mb4
    COMMENT ='题目表';

CREATE TABLE `solution`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `problem_id`   bigint(20) unsigned NOT NULL COMMENT '题目ID',
    `title`        varchar(255)        NOT NULL COMMENT '题解标题',
    `language`     varchar(16)         NOT NULL COMMENT '题解语言',
    `content`      text                NOT NULL COMMENT '题解答案',
    `caption`      varchar(2048)       NOT NULL COMMENT '题解思路',
    `created_time` datetime            NOT NULL COMMENT '创建时间',
    `updated_time` datetime            NOT NULL COMMENT '更新时间',
    `deleted_time` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`),
    KEY `idx_solution_problem_id_01` (`problem_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 5
  DEFAULT CHARSET = utf8mb4
    COMMENT ='题解表';
