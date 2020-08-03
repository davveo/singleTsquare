-- 吐槽相关
CREATE TABLE `tb_split` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` varchar(60) NOT NULL DEFAULT '' COMMENT '发布者',
  `content` text '' COMMENT '具体内容',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='吐槽表';

-- 吐槽评论
CREATE TABLE `tb_split_reply` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` varchar(60) NOT NULL DEFAULT '' COMMENT '回答用户',
  `split_id` varchar(60) NOT NULL DEFAULT '' COMMENT '吐槽id',
  `content` text '' COMMENT '具体内容',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问题表';