/*
SQLyog Ultimate v11.52 (64 bit)
MySQL - 10.1.21-MariaDB : Database - fishtimer
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
USE `fishtimer`;

/*Table structure for table `go_account` */

CREATE TABLE `go_account` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uname` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `pwd` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `createtime` bigint(20) DEFAULT '0',
  `lastLoginTime` bigint(20) DEFAULT '0' COMMENT '最后登录时间',
  `lastModifyTime` bigint(20) DEFAULT '0' COMMENT '最后修改时间',
  `nickname` varchar(30) COLLATE utf8_bin DEFAULT NULL,
  `phone` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `tel` varchar(20) COLLATE utf8_bin DEFAULT NULL COMMENT '固定电话',
  `roleID` smallint(6) NOT NULL COMMENT '角色ID',
  `superID` int(11) DEFAULT '0' COMMENT '上级ID',
  `proportions` float DEFAULT '0' COMMENT '分成比例',
  `recommend` varchar(8) COLLATE utf8_bin DEFAULT NULL COMMENT '推介码',
  `valid` smallint(6) DEFAULT '0' COMMENT '用户  1:启用 0:禁用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`uname`),
  UNIQUE KEY `unique1` (`recommend`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account` */

insert  into `go_account`(`id`,`uname`,`pwd`,`createtime`,`lastLoginTime`,`lastModifyTime`,`nickname`,`phone`,`tel`,`roleID`,`superID`,`proportions`,`recommend`,`valid`) values (8,'rjc','123',1515999939,1518102630,1517839706,'刘德华','18888888888','021000000',1,0,100,'rjc8',1),(53,'d1','123',1517828146,0,0,'代理1','111','111',4,8,5,NULL,1),(54,'d2','9527',1517836041,0,0,'代理2','123','123',4,53,10,NULL,1),(55,'d3','123',1517839430,1518102701,1518103781,'代理商3','123','123',4,8,20,NULL,1),(57,'d3_1','888888',1518088091,0,0,'动画','1','1',5,55,20,NULL,1);

/*Table structure for table `go_account_bank` */

CREATE TABLE `go_account_bank` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `accountName` varchar(50) COLLATE utf8_bin DEFAULT '',
  `accountCardNum` varchar(50) COLLATE utf8_bin DEFAULT '',
  `openBank` varchar(50) COLLATE utf8_bin DEFAULT '',
  `branchBank` varchar(100) COLLATE utf8_bin DEFAULT '',
  `balance` decimal(10,2) DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_bank` */

insert  into `go_account_bank`(`id`,`accountID`,`accountName`,`accountCardNum`,`openBank`,`branchBank`,`balance`) values (1,55,'2','2','2','2','0.00');

/*Table structure for table `go_account_child` */

CREATE TABLE `go_account_child` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `childID` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_child` */

insert  into `go_account_child`(`id`,`accountID`,`childID`) values (13,8,8),(14,8,53),(15,8,54),(16,8,55),(17,8,57),(18,53,53),(19,53,54),(20,54,54),(21,55,55),(22,55,57),(23,57,57);

/*Table structure for table `go_account_profile` */

CREATE TABLE `go_account_profile` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `avatarID` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_profile` */

insert  into `go_account_profile`(`id`,`accountID`,`avatarID`) values (1,8,1),(2,8,1);

/*Table structure for table `go_account_sale` */

CREATE TABLE `go_account_sale` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `saleamt` decimal(10,2) DEFAULT '0.00',
  `saledate` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_sale` */

insert  into `go_account_sale`(`id`,`accountID`,`saleamt`,`saledate`) values (1,11,'100.00',1517469991),(2,2,'50.00',1517469991),(3,5,'100.00',1517469991),(4,7,'200.00',1517469991),(5,4,'100.00',1517469991),(6,9,'200.00',1517469991),(7,10,'500.00',1517469991),(8,10,'200.00',1517469991),(9,53,'100.00',1517837401),(10,53,'50.00',1517837401),(11,53,'100.00',1517837401),(12,53,'200.00',1517837401),(13,54,'100.00',1517837401),(14,54,'200.00',1517837401),(15,54,'500.00',1517837401),(16,54,'200.00',1517837401);

/*Table structure for table `go_account_settlement` */

CREATE TABLE `go_account_settlement` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `superID` int(11) DEFAULT '0',
  `childID` int(11) DEFAULT '0',
  `amt` decimal(10,2) DEFAULT '0.00',
  `proportions` float DEFAULT '0',
  `settledate` bigint(20) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=553 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_settlement` */

insert  into `go_account_settlement`(`id`,`accountID`,`superID`,`childID`,`amt`,`proportions`,`settledate`) values (182,11,10,0,'20.00',20,1517476513),(183,10,0,11,'16.00',20,1517476513),(184,7,0,11,'19.20',30,1517476513),(185,2,0,11,'22.40',50,1517476513),(186,1,0,11,'22.40',100,1517476513),(187,2,1,0,'25.00',50,1517476513),(188,1,0,2,'25.00',100,1517476513),(189,5,1,0,'50.00',50,1517476513),(190,1,0,5,'50.00',100,1517476513),(191,7,2,0,'60.00',30,1517476513),(192,2,0,7,'70.00',50,1517476513),(193,1,0,7,'70.00',100,1517476513),(194,4,0,0,'50.00',50,1517476513),(195,9,4,0,'60.00',30,1517476513),(196,4,0,9,'70.00',50,1517476513),(197,10,7,0,'100.00',20,1517476513),(198,7,0,10,'120.00',30,1517476513),(199,2,0,10,'140.00',50,1517476513),(200,1,0,10,'140.00',100,1517476513),(201,10,7,0,'40.00',20,1517476513),(202,7,0,10,'48.00',30,1517476513),(203,2,0,10,'56.00',50,1517476513),(204,1,0,10,'56.00',100,1517476513),(533,53,8,0,'5.00',5,1517841638),(534,8,0,53,'95.00',100,1517841638),(535,53,8,0,'2.50',5,1517841638),(536,8,0,53,'47.50',100,1517841638),(537,53,8,0,'5.00',5,1517841638),(538,8,0,53,'95.00',100,1517841638),(539,53,8,0,'10.00',5,1517841638),(540,8,0,53,'190.00',100,1517841638),(541,54,53,0,'10.00',10,1517841638),(542,53,0,54,'4.50',5,1517841638),(543,8,0,54,'85.50',100,1517841638),(544,54,53,0,'20.00',10,1517841638),(545,53,0,54,'9.00',5,1517841638),(546,8,0,54,'171.00',100,1517841638),(547,54,53,0,'50.00',10,1517841638),(548,53,0,54,'22.50',5,1517841638),(549,8,0,54,'427.50',100,1517841638),(550,54,53,0,'20.00',10,1517841638),(551,53,0,54,'9.00',5,1517841638),(552,8,0,54,'171.00',100,1517841638);

/*Table structure for table `go_account_super` */

CREATE TABLE `go_account_super` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `accountID` int(11) DEFAULT '0',
  `superID` int(11) DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_account_super` */

insert  into `go_account_super`(`id`,`accountID`,`superID`) values (17,1,1),(18,2,2),(19,2,1),(20,3,3),(21,3,1),(22,4,4),(23,5,5),(24,5,1),(25,6,6),(26,6,1),(27,7,7),(28,7,2),(29,7,1),(30,8,8),(31,8,2),(32,8,1),(33,9,9),(34,9,4),(35,10,10),(36,10,7),(37,10,2),(38,10,1),(39,11,11),(40,11,10),(41,11,7),(42,11,2),(43,11,1);

/*Table structure for table `go_menu` */

CREATE TABLE `go_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `classID` int(11) DEFAULT NULL COMMENT '父类ID',
  `nickname` varchar(30) COLLATE utf8_bin DEFAULT NULL,
  `authurl` varchar(100) COLLATE utf8_bin NOT NULL COMMENT '菜单鉴权路由,唯一',
  `url` varchar(200) COLLATE utf8_bin DEFAULT NULL COMMENT '菜单访问路径',
  `sort` smallint(6) DEFAULT '0',
  `valid` smallint(6) DEFAULT '0' COMMENT '0:正常 1:禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_menu` */

insert  into `go_menu`(`id`,`classID`,`nickname`,`authurl`,`url`,`sort`,`valid`) values (1,5,'修改密码','/v2/account/modifypwd','/pages/account/password.html',0,1),(2,3,'仪表盘','/v2/dashboard','/index.html',0,1),(3,2,'玩家管理','/v2/player','/pages/player/index.html',10,1),(4,2,'充值管理','/v2/recharge',NULL,0,1),(5,1,'用户管理','/v2/account','/pages/account/index.html',12,1),(6,0,'添加用户','/v2/account/add',NULL,0,1),(7,0,'修改用户','/v2/account/modify',NULL,0,1),(8,0,'查看用户','/v2/account/detail','',0,1),(9,1,'主菜单','/v2/menu/menulist','/pages/menu/index.html',10,1),(10,0,'新增主菜单','/v2/menu/add',NULL,0,1),(14,0,'修改主菜单','/v2/menu/modify',NULL,0,1),(15,1,'子菜单','/v2/menu/menulistc','/pages/menu/child.html',9,1),(16,0,'新增子菜单','/v2/menu/addc',NULL,0,1),(19,0,'修改子菜单','/v2/menu/modifyc','2',0,1),(26,1,'角色','/v2/role','/pages/role/index.html',7,1),(27,0,'新增角色','/v2/role/add','',0,1),(29,0,'修改角色','/v2/role/modify','',0,1),(30,0,'角色菜单列表','/v2/role/menu','',0,1),(31,0,'修改角色菜单','/v2/role/menu/modify','',0,1),(32,0,'玩家上分','/v2/player/addpoints','',0,1),(33,5,'基本信息','','/pages/account/profile.html',10,1),(35,5,'下属账户','/v2/account','/pages/account/index.html',7,1),(36,5,'收入明细','','',9,1),(37,5,'玩家明细','','',8,1),(38,0,'角色列表接口','/v2/role','',1,1),(39,0,'账户银行信息明细','/v2/account/bank','',1,1),(40,0,'修改基本信息','/v2/account/bank/modify','',1,1);

/*Table structure for table `go_menu_class` */

CREATE TABLE `go_menu_class` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `ico` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `url` varchar(200) COLLATE utf8_bin DEFAULT '#',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_menu_class` */

insert  into `go_menu_class`(`id`,`nickname`,`ico`,`url`) values (0,'-','',''),(1,'用户管理','fa fa-files-o','/index.html'),(2,'玩家管理','fa fa-wechat','fa fa-files-o'),(3,'财务管理','fa fa-pie-chart','#'),(4,'百度','fa fa-share','http://qq.com'),(5,'仪表盘','fa fa-files-o','#');

/*Table structure for table `go_menu_role` */

CREATE TABLE `go_menu_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menuID` int(11) DEFAULT NULL,
  `roleID` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=300 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_menu_role` */

insert  into `go_menu_role`(`id`,`menuID`,`roleID`) values (6,1,2),(7,2,2),(8,3,2),(9,4,2),(10,1,3),(11,2,3),(12,3,3),(15,1,5),(16,3,6),(17,1,6),(212,5,1),(213,1,1),(214,9,1),(215,4,1),(216,3,1),(217,2,1),(218,15,1),(219,26,1),(220,31,1),(221,30,1),(222,29,1),(223,27,1),(224,19,1),(225,16,1),(226,14,1),(227,10,1),(228,8,1),(229,7,1),(230,6,1),(231,32,1),(232,33,1),(290,1,4),(291,40,4),(292,39,4),(293,37,4),(294,36,4),(295,35,4),(296,33,4),(297,38,4),(298,8,4),(299,7,4);

/*Table structure for table `go_role` */

CREATE TABLE `go_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) COLLATE utf8_bin DEFAULT NULL,
  `valid` smallint(6) DEFAULT '0' COMMENT '0:正常 1:禁用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `go_role` */

insert  into `go_role`(`id`,`nickname`,`valid`) values (1,'管理员',1),(2,'交易所',1),(3,'会员单位',1),(4,'代理商',1),(5,'经纪人',1),(6,'接班人1',1),(7,'测试权限',1);

/*Table structure for table `person` */

CREATE TABLE `person` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `first_name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `last_name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

/*Data for the table `person` */

insert  into `person`(`id`,`first_name`,`last_name`) values (120,'rjc','999'),(121,'rjc','999'),(122,'rjc1','9991'),(123,'rjc1','9991'),(124,'','');

/* Procedure structure for procedure `modifypwd` */

DELIMITER $$

/*!50003 CREATE DEFINER=`root`@`localhost` PROCEDURE `modifypwd`(
IN uid INT,
IN oldpwd VARCHAR(50),
IN newpwd VARCHAR(50)
)
BEGIN	
DECLARE _oldpwd VARCHAR(50);#原来的密码
/*只要发生异常就回滚*/ 
DECLARE EXIT HANDLER FOR SQLEXCEPTION 
BEGIN
    ROLLBACK;
  /*返回异常处理结果等其它操作*/ 
  SELECT 1;
END; 
START TRANSACTION;
	
	SELECT pwd into _oldpwd FROM go_account WHERE id=uid;
	if oldpwd=_oldpwd then 
		update go_account set pwd=newpwd where id=uid;
		select 0;
	else
		select 2;
	end if;
 
	COMMIT;
	
    END */$$
DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
