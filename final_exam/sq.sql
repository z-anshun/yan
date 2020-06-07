/*
SQLyog Professional v12.08 (64 bit)
MySQL - 5.7.19 : Database - game
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`game` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `game`;

/*Table structure for table `file_records` */

DROP TABLE IF EXISTS `file_records`;

CREATE TABLE `file_records` (
  `room_id` int(10) DEFAULT NULL,
  `end_time` int(11) DEFAULT NULL,
  `record` text
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `file_records` */

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `na_pa` (`name`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

/*Data for the table `users` */

insert  into `users`(`id`,`name`,`password`,`created_at`,`updated_at`,`deleted_at`) values (6,'as','123','2020-06-06 14:01:16','2020-06-06 14:01:16',NULL),(7,'kk','222','2020-06-07 01:36:08','2020-06-07 01:36:08',NULL),(8,'b','1','2020-06-07 06:42:27','2020-06-07 06:42:27',NULL),(9,'s','1','2020-06-07 06:44:22','2020-06-07 06:44:22',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
