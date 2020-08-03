CREATE TABLE `tb_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `zh_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '标签名中文名',
  `en_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '标签名拼音名',
  `is_hot` tinyint(5) NOT NULL  COMMENT '是否为热门标签',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签表';


CREATE TABLE `tb_city` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `zh_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '中文名',
  `en_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '拼音名',
  `country_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'country_id',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='城市';


CREATE TABLE `tb_country` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `zh_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '中文名',
  `en_name` varchar(30) NOT NULL unique DEFAULT '' COMMENT '标拼音名',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='国家';