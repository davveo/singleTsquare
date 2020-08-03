-- 问题相关
CREATE TABLE `question` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` varchar(60) NOT NULL DEFAULT '' COMMENT '问题发布者',
  `title` varchar(60) NOT NULL DEFAULT '' COMMENT '问题标题',
  `desc` varchar(60) NOT NULL DEFAULT '' COMMENT '问题描述',
  `content` text '' COMMENT '具体内容',
  `title` varchar(60) NOT NULL DEFAULT '' COMMENT '问题标题',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问题表';

-- 问题关联的标签
CREATE TABLE `question_tag` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `tag_id` varchar(60) NOT NULL DEFAULT '' COMMENT '标签id',
  `question_id` varchar(60) NOT NULL DEFAULT '' COMMENT '问题id',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问题关联标签';

-- 回答相关
CREATE TABLE `answer` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_id` varchar(60) NOT NULL DEFAULT '' COMMENT '回答用户',
  `question_id` varchar(60) NOT NULL DEFAULT '' COMMENT '回答问题',
  `content` text '' COMMENT '具体内容',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='问题表';