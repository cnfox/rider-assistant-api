CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `openid` int NOT NULL COMMENT '微信id',
  `name` varchar(255) DEFAULT NULL COMMENT '江湖称号',
  `nickname` varchar(255) DEFAULT NULL COMMENT '昵称',
  `avatar` varchar(255) DEFAULT NULL COMMENT '头像',
  `group` int DEFAULT '0' COMMENT '分组',
  `score` int DEFAULT '0' COMMENT '积分',
  `is_leader` tinyint DEFAULT '0' COMMENT '是否领队:1是、0否',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='队员表';

CREATE TABLE `event` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '活动标题',
  `leader_id` int NOT NULL COMMENT '领队id',
  `leader_name` varchar(255) NOT NULL COMMENT '领队花名',
  `event_date` varchar(255) NOT NULL COMMENT '活动日期',
  `start_time` varchar(255) DEFAULT NULL COMMENT '集合时间',
  `event_route` varchar(255) DEFAULT NULL COMMENT '活动路线',
  `strength` varchar(255) DEFAULT NULL COMMENT '活动强度',
  `estimate_time` varchar(255) DEFAULT NULL COMMENT '预计耗时',
  `notice` varchar(255) DEFAULT NULL COMMENT '特别说明',
  `code` varchar(255) NOT NULL COMMENT '签到口令',
  `score` int DEFAULT '1' COMMENT '签到积分',
  `upper_limit` int DEFAULT '999' COMMENT '人数上限',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='活动表';

CREATE TABLE `event_group` (
  `id` int NOT NULL AUTO_INCREMENT,
  `event_id` int NOT NULL COMMENT '活动id',
  `sequence` varchar(255) NOT NULL COMMENT '分组序号',
  `leader_id` int NOT NULL COMMENT '领队id',
  `leader_name` varchar(255) NOT NULL COMMENT '领队花名',
  `strength` varchar(255) DEFAULT NULL COMMENT '分组强度',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='活动分组表';

CREATE TABLE `event_enroll` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '用户id',
  `event_id` int NOT NULL COMMENT '活动id',
  `check_in` tinyint DEFAULT '0' COMMENT '是否签到1是0否',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='活动报名签到表';

CREATE TABLE `score_record` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int NOT NULL COMMENT '队员id',
  `type` tinyint NOT NULL COMMENT '加积分原因:1活动签到、2历史积分、3队服、4全程、5文章、6视频',
  `score` int NOT NULL COMMENT '积分',
  `note` varchar(255) DEFAULT NULL COMMENT '备注',
  `event_id` int DEFAULT '0' COMMENT '活动id',
  `admin_id` int DEFAULT '0' COMMENT '管理员id',
  `create_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB COMMENT='积分流水表';