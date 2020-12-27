/*
SQLyog Ultimate
MySQL - 5.5.22 : Database - alterra_store
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`alterra_store` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `alterra_store`;

/*Table structure for table `cart` */

DROP TABLE IF EXISTS `cart`;

CREATE TABLE `cart` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` int(11) DEFAULT NULL,
  `checkout` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

/*Data for the table `cart` */

insert  into `cart`(`id`,`product_id`,`checkout`) values (3,2,1);
insert  into `cart`(`id`,`product_id`,`checkout`) values (7,3,1);

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

/*Data for the table `category` */

insert  into `category`(`id`,`name`) values (1,'Stationary');
insert  into `category`(`id`,`name`) values (2,'Phone Acc');

/*Table structure for table `product` */

DROP TABLE IF EXISTS `product`;

CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `product_ibfk_1` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1;

/*Data for the table `product` */

insert  into `product`(`id`,`name`,`category_id`) values (1,'Pencil',1);
insert  into `product`(`id`,`name`,`category_id`) values (2,'Pen',1);
insert  into `product`(`id`,`name`,`category_id`) values (3,'Eraser',1);
insert  into `product`(`id`,`name`,`category_id`) values (4,'Ruler',1);
insert  into `product`(`id`,`name`,`category_id`) values (5,'Phone Charger',2);
insert  into `product`(`id`,`name`,`category_id`) values (6,'Screen Protector',2);
insert  into `product`(`id`,`name`,`category_id`) values (7,'Head Set',2);
insert  into `product`(`id`,`name`,`category_id`) values (8,'Ear Phone',2);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
