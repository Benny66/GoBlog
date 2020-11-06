/*
Navicat MySQL Data Transfer

Source Server         : 127.0.0.1
Source Server Version : 50505
Source Host           : localhost:3306
Source Database       : goblog

Target Server Type    : MYSQL
Target Server Version : 50505
File Encoding         : 65001

Date: 2020-11-06 17:31:04
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `blog_articles`
-- ----------------------------
DROP TABLE IF EXISTS `blog_articles`;
CREATE TABLE `blog_articles` (
  `article_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `user_id` int(10) NOT NULL COMMENT '用户id',
  `title` varchar(200) NOT NULL COMMENT '文章标题',
  `content` longtext NOT NULL COMMENT '文章内容',
  `status` varchar(16) NOT NULL DEFAULT 'publish' COMMENT '状态:public=公开,private=私有',
  `views` int(10) NOT NULL DEFAULT '0' COMMENT '浏览量',
  `like_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`article_id`),
  KEY `idx_blog_articles_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_articles
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_comments`
-- ----------------------------
DROP TABLE IF EXISTS `blog_comments`;
CREATE TABLE `blog_comments` (
  `comment_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `user_id` int(10) NOT NULL COMMENT '发表用户ID',
  `article_id` int(10) NOT NULL COMMENT '评论博文ID',
  `like_count` bigint(20) NOT NULL COMMENT '点赞数',
  `date` datetime DEFAULT NULL COMMENT '评论日期',
  `content` text NOT NULL COMMENT '评论内容',
  `p_comment_id` bigint(20) NOT NULL COMMENT '父评论ID',
  PRIMARY KEY (`comment_id`),
  KEY `article_id` (`article_id`) USING BTREE,
  KEY `date` (`date`) USING BTREE,
  KEY `p_comment_id` (`p_comment_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_comments
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_labels`
-- ----------------------------
DROP TABLE IF EXISTS `blog_labels`;
CREATE TABLE `blog_labels` (
  `label_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `name` varchar(20) NOT NULL COMMENT '标签名称',
  `alias` varchar(15) NOT NULL COMMENT '标签别名',
  `description` text NOT NULL COMMENT '标签描述',
  PRIMARY KEY (`label_id`),
  KEY `name` (`name`) USING BTREE,
  KEY `alias` (`alias`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_labels
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_options`
-- ----------------------------
DROP TABLE IF EXISTS `blog_options`;
CREATE TABLE `blog_options` (
  `name` varchar(32) NOT NULL COMMENT '配置名',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '配置用户id',
  `value` text COMMENT '配置内容',
  PRIMARY KEY (`name`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_options
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_set_article_label`
-- ----------------------------
DROP TABLE IF EXISTS `blog_set_article_label`;
CREATE TABLE `blog_set_article_label` (
  `article_id` int(10) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `label_id` int(10) NOT NULL COMMENT '标签ID',
  PRIMARY KEY (`article_id`),
  KEY `label_id` (`label_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_set_article_label
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_set_article_sort`
-- ----------------------------
DROP TABLE IF EXISTS `blog_set_article_sort`;
CREATE TABLE `blog_set_article_sort` (
  `article_id` int(10) NOT NULL COMMENT '文章ID',
  `sort_id` int(10) NOT NULL COMMENT '分类ID',
  PRIMARY KEY (`article_id`,`sort_id`),
  KEY `sort_id` (`sort_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_set_article_sort
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_set_users_platform`
-- ----------------------------
DROP TABLE IF EXISTS `blog_set_users_platform`;
CREATE TABLE `blog_set_users_platform` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int(10) NOT NULL COMMENT '用户id',
  `type` varchar(10) NOT NULL COMMENT '第三方平台类型',
  `openid` varchar(64) NOT NULL COMMENT '第三方唯一标识',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_set_users_platform
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_sorts`
-- ----------------------------
DROP TABLE IF EXISTS `blog_sorts`;
CREATE TABLE `blog_sorts` (
  `sort_id` int(10) NOT NULL COMMENT '分类ID',
  `name` varchar(50) NOT NULL COMMENT '分类名称',
  `alias` varchar(15) NOT NULL COMMENT '分类别名',
  `description` text NOT NULL COMMENT '分类描述',
  `p_sort_id` bigint(20) NOT NULL COMMENT '父分类ID',
  PRIMARY KEY (`sort_id`),
  KEY `name` (`name`) USING BTREE,
  KEY `alias` (`alias`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_sorts
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_users`
-- ----------------------------
DROP TABLE IF EXISTS `blog_users`;
CREATE TABLE `blog_users` (
  `user_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(32) NOT NULL COMMENT '登录账号',
  `password` varchar(64) NOT NULL COMMENT '登陆密码',
  `email` varchar(100) DEFAULT NULL COMMENT '用户邮箱',
  `avatar_url` varchar(200) DEFAULT NULL COMMENT '用户头像',
  `screen_name` varchar(32) NOT NULL COMMENT '昵称',
  `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_time` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `name` (`name`) USING BTREE,
  UNIQUE KEY `mail` (`email`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_users
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_users_qq`
-- ----------------------------
DROP TABLE IF EXISTS `blog_users_qq`;
CREATE TABLE `blog_users_qq` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int(10) NOT NULL COMMENT '用户id',
  `openid` varchar(64) NOT NULL COMMENT '用户唯一标识',
  `nick_name` varchar(100) NOT NULL COMMENT 'qq昵称',
  `avatar_url` varchar(200) NOT NULL COMMENT 'qq头像',
  `city` varchar(50) NOT NULL COMMENT '城市',
  `province` varchar(50) NOT NULL COMMENT '省份',
  `country` varchar(50) NOT NULL COMMENT '国家',
  `gender` varchar(10) NOT NULL COMMENT '性别',
  `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_time` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_qq_openid` (`openid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_users_qq
-- ----------------------------

-- ----------------------------
-- Table structure for `blog_users_wx`
-- ----------------------------
DROP TABLE IF EXISTS `blog_users_wx`;
CREATE TABLE `blog_users_wx` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int(10) NOT NULL COMMENT '用户id',
  `openid` varchar(64) NOT NULL COMMENT '用户唯一标识',
  `nick_name` varchar(100) NOT NULL COMMENT '微信昵称',
  `avatar_url` varchar(200) NOT NULL COMMENT '微信头像',
  `city` varchar(50) NOT NULL COMMENT '城市',
  `province` varchar(50) NOT NULL COMMENT '省份',
  `country` varchar(50) NOT NULL COMMENT '国家',
  `gender` varchar(10) NOT NULL COMMENT '性别',
  `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_time` datetime DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_wx_openid` (`openid`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of blog_users_wx
-- ----------------------------