/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : localhost:3306
 Source Schema         : app_admin

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 06/01/2022 22:09:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_pay
-- ----------------------------
DROP TABLE IF EXISTS `app_pay`;
CREATE TABLE `app_pay` (
  `users_id` int(10) NOT NULL COMMENT '用户id',
  `pay_id` bigint(20) NOT NULL COMMENT '支付订单号',
  `pay_time` datetime(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '支付时间',
  `pay_complete` tinyint(1) DEFAULT NULL COMMENT '是否完成支付',
  `pay_money` varchar(10) DEFAULT NULL COMMENT '支付金额',
  `pay_vip` tinyint(1) DEFAULT NULL COMMENT '是否是vip',
  `pay_client` varchar(20) DEFAULT NULL COMMENT '支付客户端',
  `pay_source` varchar(30) DEFAULT NULL COMMENT '支付来源',
  `pay_intent` varchar(30) DEFAULT 'pay' COMMENT '支付意图',
  PRIMARY KEY (`pay_id`) USING BTREE,
  UNIQUE KEY `indexPAY` (`pay_id`) USING HASH COMMENT '索引单号'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_pay
-- ----------------------------
BEGIN;
INSERT INTO `app_pay` VALUES (100001, 1000001, '2022-01-04 22:52:22.296331', NULL, NULL, NULL, NULL, NULL, 'pay');
COMMIT;

-- ----------------------------
-- Table structure for app_users
-- ----------------------------
DROP TABLE IF EXISTS `app_users`;
CREATE TABLE `app_users` (
  `users_id` int(10) NOT NULL COMMENT '用户ID',
  `users_nickname` varchar(30) DEFAULT NULL COMMENT '用户昵称',
  `users_password` varchar(255) DEFAULT NULL COMMENT '用户密码',
  `users_phone` varchar(11) DEFAULT NULL COMMENT '用户手机号',
  `users_create` datetime(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '用户创建时间',
  `users_ portrait` varchar(255) DEFAULT NULL COMMENT '用户头像',
  `users_client` varchar(20) DEFAULT NULL COMMENT '用户注册端',
  `users_source` varchar(30) DEFAULT NULL COMMENT '用户来源',
  `user_close` tinyint(1) DEFAULT '1' COMMENT '是否关闭用户',
  PRIMARY KEY (`users_id`) USING BTREE,
  UNIQUE KEY `indexUSERS` (`users_id`) USING HASH COMMENT '主键唯一索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_users
-- ----------------------------
BEGIN;
INSERT INTO `app_users` VALUES (100000, '帅威', '123456', '13151289178', '2022-01-04 22:07:20.125690', NULL, 'myb_android', NULL, 1);
INSERT INTO `app_users` VALUES (1000001, 'sw', 'qqq', '13151289178', '2022-01-04 22:08:35.356836', NULL, 'hmy_android', NULL, 1);
COMMIT;

-- ----------------------------
-- Table structure for app_vip
-- ----------------------------
DROP TABLE IF EXISTS `app_vip`;
CREATE TABLE `app_vip` (
  `users_id` int(10) NOT NULL COMMENT '用户ID',
  `pay_id` bigint(20) DEFAULT NULL COMMENT '支付单号',
  `vip_create` datetime(6) DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(6) COMMENT 'VIP 开始时间',
  `vip_end` datetime(6) DEFAULT NULL COMMENT 'VIP 结束时间',
  PRIMARY KEY (`users_id`),
  UNIQUE KEY `indexVIP` (`users_id`) USING HASH COMMENT 'VIP索引'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of app_vip
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
