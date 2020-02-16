CREATE TABLE `todoList`.`tasks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `Account` varchar(255) NOT NULL DEFAULT '',
  `Name ` varchar(255) NOT NULL DEFAULT '',
  `Paaswd` varchar(255) DEFAULT '' COMMENT 'titleだけでもOK',
  `Created` datetime DEFAULT NULL
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
