/*
 Navicat MySQL Data Transfer

 Source Server         : localhost mysql dbngin
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : hassandb

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 05/12/2023 18:53:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for schema_migrations
-- ----------------------------
DROP TABLE IF EXISTS `schema_migrations`;
CREATE TABLE `schema_migrations` (
  `version` varchar(255) CHARACTER SET latin1 COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

-- ----------------------------
-- Records of schema_migrations
-- ----------------------------
BEGIN;
INSERT INTO `schema_migrations` VALUES ('20211210045538');
COMMIT;

-- ----------------------------
-- Table structure for todos
-- ----------------------------
DROP TABLE IF EXISTS `todos`;
CREATE TABLE `todos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `completed` int DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `token` text,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
