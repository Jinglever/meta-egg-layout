SET FOREIGN_KEY_CHECKS=0;
CREATE DATABASE `meta-egg-layout` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
USE `meta-egg-layout`;
CREATE TABLE `user` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(64) COMMENT '用户名',
  `gender` BIGINT UNSIGNED NOT NULL COMMENT '性别',
  `age` TINYINT UNSIGNED NOT NULL COMMENT '年龄',
  `is_on_job` TINYINT NOT NULL DEFAULT '0' COMMENT '是否在职',
  `birthday` DATE COMMENT '生日',
  `created_by` BIGINT UNSIGNED COMMENT '创建者',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `updated_by` BIGINT UNSIGNED COMMENT '更新者',
  `updated_at` DATETIME NOT NULL COMMENT '更新时间',
  `deleted_by` BIGINT UNSIGNED COMMENT '删除者',
  `deleted_at` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  UNIQUE (`name`),
  INDEX (`deleted_at`),
  INDEX (`is_on_job`, `age`),
  FOREIGN KEY (`gender`) REFERENCES `gender` (`id`),
  FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  FOREIGN KEY (`updated_by`) REFERENCES `user` (`id`),
  FOREIGN KEY (`deleted_by`) REFERENCES `user` (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户';
CREATE TABLE `gender` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `sematic` VARCHAR(8) NOT NULL COMMENT '语义',
  `desc` VARCHAR(64) COMMENT '描述',
  `deleted_at` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  INDEX (`deleted_at`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='性别';
CREATE TABLE `user_phone` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `phone_number` VARCHAR(20) NOT NULL COMMENT '电话号码',
  `created_by` BIGINT UNSIGNED COMMENT '创建者',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `deleted_by` BIGINT UNSIGNED COMMENT '删除者',
  `deleted_at` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  INDEX (`deleted_at`),
  UNIQUE (`phone_number`, `deleted_at`),
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  FOREIGN KEY (`deleted_by`) REFERENCES `user` (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户电话';
CREATE TABLE `user_tag` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户ID',
  `tag_id` BIGINT UNSIGNED NOT NULL COMMENT '标签ID',
  `created_by` BIGINT UNSIGNED COMMENT '创建者',
  `created_at` DATETIME NOT NULL COMMENT '创建时间',
  `deleted_by` BIGINT UNSIGNED COMMENT '删除者',
  `deleted_at` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  INDEX (`deleted_at`),
  UNIQUE (`user_id`, `tag_id`, `deleted_at`),
  FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`),
  FOREIGN KEY (`created_by`) REFERENCES `user` (`id`),
  FOREIGN KEY (`deleted_by`) REFERENCES `user` (`id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='用户标签关系';
CREATE TABLE `tag` (
  `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(64) NOT NULL COMMENT '标签名称',
  `deleted_at` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  INDEX (`deleted_at`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='标签';
