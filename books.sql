/*
 Navicat Premium Data Transfer

 Source Server         : MacMysql
 Source Server Type    : MySQL
 Source Server Version : 50731
 Source Host           : localhost:3306
 Source Schema         : booksystem

 Target Server Type    : MySQL
 Target Server Version : 50731
 File Encoding         : 65001

 Date: 24/12/2020 10:29:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `title` varchar(20) DEFAULT NULL,
  `author` varchar(10) DEFAULT NULL,
  `price` float(10,2) DEFAULT NULL,
  `stock` int(10) DEFAULT NULL COMMENT '库存',
  `imgpth` varchar(255) DEFAULT NULL COMMENT '图书封面',
  `sales` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of books
-- ----------------------------
BEGIN;
INSERT INTO `books` VALUES (1, '解忧杂货店', '东野圭吾', 27.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (2, '边城', '沈从文', 23.00, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (3, '中国哲学史', '冯友兰', 44.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (4, '忽然七日', ' 劳伦', 19.33, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (5, '苏东坡传', '林语堂', 19.30, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (6, '百年孤独', '马尔克斯', 29.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (7, '扶桑', '严歌苓', 19.80, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (8, '给孩子的诗', '北岛', 22.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (9, '为奴十二年', '所罗门', 16.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (10, '平凡的世界', '路遥', 55.00, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (11, '悟空传', '今何在', 14.00, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (12, '硬派健身', '斌卡', 31.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (13, '从晚清到民国', '唐德刚', 39.90, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (14, '三体', '刘慈欣', 56.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (15, '看见', '柴静', 19.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (16, '活着', '余华', 11.00, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (17, '小王子', '安托万', 19.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (18, '我们仨', '杨绛', 11.30, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (19, '生命不息,折腾不止', '罗永浩', 25.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (20, '皮囊', '蔡崇达', 23.90, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (21, '恰到好处的幸福', '毕淑敏', 16.40, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (22, '大数据预测', '埃里克', 37.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (23, '人月神话', '布鲁克斯', 55.90, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (24, 'C语言入门经典', '霍尔顿', 45.00, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (25, '数学之美', '吴军', 29.90, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (26, 'Java编程思想', '埃史尔', 70.50, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (27, '设计模式之禅', '秦小波', 20.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (28, '图解机器学习', '杉山将', 33.80, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (29, '艾伦图灵传', '安德鲁', 47.20, 100, 'static/img/default.jpg', 100);
INSERT INTO `books` VALUES (30, '教父', '马里奥普佐', 29.00, 100, 'static/img/default.jpg', 100);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
