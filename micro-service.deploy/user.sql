-- Adminer 4.2.2 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `rushteam` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_bin */;
USE `rushteam`;

DROP TABLE IF EXISTS `admin_group`;
CREATE TABLE `admin_group` (
  `gid` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_bin NOT NULL COMMENT '组名',
  PRIMARY KEY (`gid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `admin_group` (`gid`, `name`) VALUES
(1,	'超级管理员'),
(2,	'管理员');

DROP TABLE IF EXISTS `admin_policy`;
CREATE TABLE `admin_policy` (
  `id` tinyint(4) NOT NULL AUTO_INCREMENT,
  `subject` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '访问实体',
  `object` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '访问资源',
  `action` varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT '访问方法',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `admin_policy` (`id`, `subject`, `object`, `action`) VALUES
(1,	'gid_1',	'/admin/users',	'read');

DROP TABLE IF EXISTS `admin_user_group`;
CREATE TABLE `admin_user_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL,
  `gid` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;


DROP TABLE IF EXISTS `user_login`;
CREATE TABLE `user_login` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uid` bigint(20) NOT NULL COMMENT 'uid',
  `platform` char(16) COLLATE utf8mb4_bin NOT NULL DEFAULT 'username' COMMENT '登录方式@password本地登陆@wechat:微信',
  `openid` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT '登录id',
  `verified` char(1) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '是否验证@0未验证@1已验证',
  `access_token` varchar(128) COLLATE utf8mb4_bin NOT NULL COMMENT '登录秘钥',
  `access_expire` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '有效期',
  PRIMARY KEY (`id`),
  UNIQUE KEY `platform_openid` (`platform`,`openid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

INSERT INTO `user_login` (`id`, `uid`, `platform`, `openid`, `verified`, `access_token`, `access_expire`) VALUES
(1,	1,	'phone',	'18310497688',	'1',	'fb469d7ef430b0baf0cab6c436e70375',	'2018-12-12 00:00:00');

DROP TABLE IF EXISTS `user_user`;
CREATE TABLE `user_user` (
  `uid` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `gender` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT '性别@1:男@2:女',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '头像',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';

INSERT INTO `user_user` (`uid`, `nickname`, `gender`, `avatar`, `updated_at`, `created_at`) VALUES
(1,	'落舞者',	'1',	'',	'2018-11-28 16:04:27',	'2018-11-28 16:04:27');

-- 2018-11-28 18:01:47