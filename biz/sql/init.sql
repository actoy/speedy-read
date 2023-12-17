CREATE TABLE `stock_symbols` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `symbol` varchar(63) NOT NULL DEFAULT '' COMMENT '股票代码',
     `desc_zh` varchar(63) NOT NULL DEFAULT '' COMMENT '中文名称',
     `desc_en` varchar(63) NOT NULL DEFAULT '' COMMENT '英文名称',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `symbol` (`symbol`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='stock code';

CREATE TABLE `sites` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `source_id` bigint(20) unsigned NOT NULL COMMENT '关联sourceID',
     `source_type` varchar(255) NOT NULL DEFAULT ''COMMENT '关联sourceType',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '地址',
     `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `meta` (`source_id`, `source_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='sites';

CREATE TABLE `authors` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '地址',
     `author_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'author name',
     `image` varchar(1024) NOT NULL DEFAULT '' COMMENT '作者头像',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='author info';

CREATE TABLE `articles` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `author` bigint(20) unsigned NOT NULL COMMENT 'author info',
     `publish_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'publish time',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '文章链接',
     `type` varchar(20) NOT NULL DEFAULT '' COMMENT '文章类别',
     `title` text NOT NULL COMMENT '标题',
     `content` longtext NOT NULL COMMENT '正文内容',
     `status` tinyint(4) NOT NULL COMMENT 'status，1:init, 2:reject, 3:pass',
     `score` tinyint(4) NOT NULL COMMENT '分数',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     `updated_by` bigint(20) unsigned NOT NULL COMMENT 'update user',
     PRIMARY KEY (`id`),
     KEY `publish_at` (`publish_at`),
     KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article info';

CREATE TABLE `article_metas` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `article_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `meta_key` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta key',
     `meta_value` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta value',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `article_id` (`article_id`),
     KEY `meta` (`meta_key`, `meta_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article meta';

CREATE TABLE `article_summarys` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `article_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `title` text NOT NULL COMMENT '标题',
     `summary` text NOT NULL COMMENT '摘要',
     `content` longtext NOT NULL COMMENT '原文内容',
     `content_summary` text NOT NULL COMMENT '一句话原文',
     `outline` text NOT NULL COMMENT '提纲',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     `created_by` bigint(20) unsigned NOT NULL COMMENT 'create user',
     `updated_by` bigint(20) unsigned NOT NULL COMMENT 'update user',
     PRIMARY KEY (`id`),
     KEY `article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article summary';

CREATE TABLE `labels` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='labels 标签列'
