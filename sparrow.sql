-- MySQL dump 10.13  Distrib 5.7.30, for Linux (x86_64)
--
-- Host: 192.168.19.253    Database: sparrow
-- ------------------------------------------------------
-- Server version	5.7.24

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `sp_article`
--

DROP TABLE IF EXISTS `sp_article`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `ar_title` varchar(100) NOT NULL COMMENT '文章标题',
  `ar_cate` tinyint(3) unsigned NOT NULL COMMENT '文章分类，关联sp_resp_cate',
  `ar_content` text NOT NULL COMMENT '文章内容',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `userid` int(10) unsigned NOT NULL COMMENT '用户id关联sp_user',
  `liked` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞次数',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1是正常0删除',
  `ar_pure_content` text NOT NULL COMMENT '文章内容-纯文本',
  `repo_unique_code` varchar(10) NOT NULL COMMENT '知识库唯一标识',
  `ar_unique_code` varchar(10) NOT NULL COMMENT 'wiki唯一标识符',
  `private` enum('1','0') NOT NULL DEFAULT '1' COMMENT '是否是公开的1是0否',
  PRIMARY KEY (`id`),
  FULLTEXT KEY `sp_article_ar_pure_content_IDX` (`ar_pure_content`,`ar_title`),
  FULLTEXT KEY `sp_article_ar_title_IDX` (`ar_title`,`ar_pure_content`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COMMENT='文章表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_article`
--

--
-- Table structure for table `sp_book`
--

DROP TABLE IF EXISTS `sp_book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_book` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '图书id',
  `b_name` varchar(100) NOT NULL COMMENT '图书标题',
  `b_url` varchar(100) NOT NULL COMMENT '文件地址',
  `b_cate_id` tinyint(3) unsigned NOT NULL COMMENT '图书分类 关联sp_book_cate',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `b_userid` bigint(20) unsigned NOT NULL COMMENT '操作人id 关联sp_user表',
  `b_avator_url` varchar(100) NOT NULL COMMENT ' 图书封面地址',
  `download` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '下载次数',
  `book_md5` varchar(32) NOT NULL COMMENT '文件的md5值',
  `avator_md5` varchar(32) NOT NULL COMMENT '图片的md5',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='图书馆表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_book`
--

--
-- Table structure for table `sp_book_cate`
--

DROP TABLE IF EXISTS `sp_book_cate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_book_cate` (
  `id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `cate_name` varchar(100) NOT NULL COMMENT '分类名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='图书分类表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_book_cate`
--

LOCK TABLES `sp_book_cate` WRITE;
/*!40000 ALTER TABLE `sp_book_cate` DISABLE KEYS */;
INSERT INTO `sp_book_cate` VALUES (1,'软件开发'),(2,'数据库');
/*!40000 ALTER TABLE `sp_book_cate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_collect`
--

DROP TABLE IF EXISTS `sp_collect`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_collect` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `resource_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '资源类型:1是文档2：资料 3:团队 4:知识库',
  `userid` int(10) unsigned NOT NULL COMMENT '用户id关联sp_user表ID',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `resource_unique_code` varchar(10) NOT NULL COMMENT '资源类型',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='收藏表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_collect`
--

LOCK TABLES `sp_collect` WRITE;
/*!40000 ALTER TABLE `sp_collect` DISABLE KEYS */;
/*!40000 ALTER TABLE `sp_collect` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_doc`
--

DROP TABLE IF EXISTS `sp_doc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_doc` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `doc_md5` varchar(32) NOT NULL COMMENT '关联sp_files表的md5',
  `doc_name` varchar(100) NOT NULL COMMENT '资料名称',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deltime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `state` enum('0','1') NOT NULL DEFAULT '0' COMMENT '0删除1正常',
  `repo_unique_code` varchar(10) NOT NULL COMMENT '知识库唯一标识符',
  `file_dir_level` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '资料目录级别',
  `userid` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='资料表，所有的资料库汇总表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_doc`
--

--
-- Table structure for table `sp_filecate`
--

DROP TABLE IF EXISTS `sp_filecate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_filecate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT '目录名称',
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父id，关联id主键',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `repo_unique_code` varchar(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='文件分类=路径';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_filecate`
--

--
-- Table structure for table `sp_files`
--

DROP TABLE IF EXISTS `sp_files`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_files` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '资料id',
  `filename` varchar(100) NOT NULL COMMENT '文件名称',
  `filetype` varchar(10) NOT NULL COMMENT '文件类型',
  `md5` varchar(32) NOT NULL COMMENT '文件md5',
  `fileurl` varchar(100) NOT NULL COMMENT '文件地址',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `ip` varchar(15) NOT NULL DEFAULT '0' COMMENT '存储的ip地址',
  `domain` varchar(100) NOT NULL DEFAULT '0' COMMENT '域名地址',
  `userid` int(10) unsigned NOT NULL COMMENT '用户id关联sp_user表ID',
  `size` varchar(20) NOT NULL COMMENT '文件大小',
  `file_dir_level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件目录，关联sp_filecate表id，0代表是第一目录',
  `repo_unique_code` varchar(10) NOT NULL COMMENT '知识库id关联sp_repository表repo_unique_code;avator是头像信息;contentImg是文章图片信息',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sp_files_md5_IDX` (`md5`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COMMENT='资料表基准表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_files`
--

--
-- Table structure for table `sp_follow`
--

DROP TABLE IF EXISTS `sp_follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_follow` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(10) unsigned NOT NULL COMMENT '用户id关联sp_user表id',
  `follow_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '关注类型0是人 1是知识库',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `follow_id` varchar(10) NOT NULL COMMENT '关注内容id',
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sp_follow_userid_IDX` (`userid`,`follow_type`,`follow_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户关注表-关注人或者知识库';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_follow`
--

--
-- Table structure for table `sp_good_history`
--

DROP TABLE IF EXISTS `sp_good_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_good_history` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(10) unsigned NOT NULL COMMENT '用户id',
  `aritcle_unique_code` varchar(10) DEFAULT NULL COMMENT '文章的唯一标识符,关联sp_article表',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户点赞表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_good_history`
--

--
-- Table structure for table `sp_recycle_bin`
--

DROP TABLE IF EXISTS `sp_recycle_bin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_recycle_bin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `resource_unqiue_code` varchar(10) NOT NULL COMMENT '资源唯一标识符',
  `repo_cate` tinyint(3) unsigned NOT NULL COMMENT '资源类型，关联sp_resp_cate表ID',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `userid` int(10) unsigned NOT NULL COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='回收站';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_recycle_bin`
--

--
-- Table structure for table `sp_repo_cate`
--

DROP TABLE IF EXISTS `sp_repo_cate`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_repo_cate` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '知识库分类表',
  `catename` varchar(100) NOT NULL COMMENT '分类名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='知识库分类';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_repo_cate`
--

LOCK TABLES `sp_repo_cate` WRITE;
/*!40000 ALTER TABLE `sp_repo_cate` DISABLE KEYS */;
INSERT INTO `sp_repo_cate` VALUES (1,'文档知识库'),(2,'资料知识库');
/*!40000 ALTER TABLE `sp_repo_cate` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_repository`
--

DROP TABLE IF EXISTS `sp_repository`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_repository` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '知识库id主键',
  `repo_name` varchar(100) NOT NULL COMMENT '知识库名称',
  `repo_desc` varchar(255) NOT NULL DEFAULT '空' COMMENT '知识库简介',
  `repo_cate` smallint(5) unsigned NOT NULL COMMENT '知识库类型（文档、资料）关联sp_resp_cate',
  `repo_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否公开  1是 0 否',
  `repo_user_group` int(10) unsigned NOT NULL DEFAULT '9999' COMMENT '资源所属组关联sp_user_group表id，默认9999表示不属于组',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更改时间',
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `userid` bigint(20) unsigned NOT NULL COMMENT '用户id关联sp_user表id',
  `repo_unique_code` varchar(50) NOT NULL COMMENT '知识库唯一字符串',
  `state` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1是正常 0是删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sp_respository_UN` (`repo_unique_code`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COMMENT=' 个人知识库表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_repository`
--
--
-- Table structure for table `sp_user`
--

DROP TABLE IF EXISTS `sp_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id，自增',
  `username` varchar(100) NOT NULL COMMENT '用户名称',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT ' 插入时间',
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `avatar_md5` varchar(32) NOT NULL DEFAULT 'null' COMMENT '头像md5值关联sp_files的md5',
  `profile` varchar(255) NOT NULL DEFAULT '他什么也没有说' COMMENT '个人介绍',
  `phone` varchar(50) NOT NULL DEFAULT 'null' COMMENT '联系方式',
  `gold` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否为上帝 0否 1是',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_user`
--

LOCK TABLES `sp_user` WRITE;
/*!40000 ALTER TABLE `sp_user` DISABLE KEYS */;
INSERT INTO `sp_user` VALUES (1,'admin','$2a$10$Q.2q63oHQuFk0hgaarc9BOc6Y6F62J6jeQsOBOGdXMCB3Vw3LgZBq','2021-05-30 03:08:42','2021-05-30 03:08:42','2021-05-30 03:08:42','10f0d19f4c4e371fdbf78b1fdc0ca1e8','我是超级管理员啊','null',0);
/*!40000 ALTER TABLE `sp_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sp_user_group`
--

DROP TABLE IF EXISTS `sp_user_group`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_user_group` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '协作团队id',
  `group_unique_code` varchar(20) NOT NULL COMMENT '团队标识符，唯一',
  `group_name` varchar(100) NOT NULL COMMENT '团队名称',
  `group_member` varchar(100) NOT NULL DEFAULT '[]' COMMENT '组内成员列表',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletetime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `group_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否公开1公开0不公开',
  `group_desc` varchar(100) NOT NULL DEFAULT '空' COMMENT '团队描述信息',
  `group_avator_url_md5` varchar(100) NOT NULL DEFAULT '0' COMMENT '团队图片地址,关联sp_img中的id',
  `userid` int(10) unsigned NOT NULL COMMENT '该团队负责人',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sp_user_group_UN` (`group_unique_code`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COMMENT='用户协作团队表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_user_group`
--

--
-- Table structure for table `sp_user_group_member`
--

DROP TABLE IF EXISTS `sp_user_group_member`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sp_user_group_member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `userid` int(10) unsigned NOT NULL COMMENT '改组的成员id关联用户表sp_user ID',
  `createtime` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `modifytime` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `group_unique_code` varchar(20) DEFAULT NULL COMMENT '组的唯一标识符',
  `isleader` enum('1','0') NOT NULL DEFAULT '0' COMMENT '是否为团队leader 1是0否',
  PRIMARY KEY (`id`),
  UNIQUE KEY `sp_user_group_member_UN` (`userid`,`group_unique_code`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='组成员信息';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sp_user_group_member`
--
