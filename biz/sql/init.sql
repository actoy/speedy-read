CREATE TABLE `site_metas` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `site_id` bigint(20) unsigned NOT NULL COMMENT 'site_id',
     `meta_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta type',
     `meta_key` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta key',
     `meta_value` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta value',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `site_id` (`site_id`),
     KEY `meta` (`meta_key`, `meta_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='site meta';

CREATE TABLE `sites` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '地址',
     `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '简介',
     `tag` varchar(63) NOT NULL DEFAULT ''COMMENT '网站标识',
     `type` varchar(20) NOT NULL DEFAULT '' COMMENT '网站类型',
     `type_key` varchar(128) NOT NULL DEFAULT '' COMMENT '不同类型网站对应的key，如八抓鱼的taskID',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='sites';

CREATE TABLE `authors` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '地址',
     `author_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'author name',
     `image` varchar(1024) NOT NULL DEFAULT '' COMMENT '作者头像',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='author info';

CREATE TABLE `articles` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `author_id` bigint(20) unsigned NOT NULL COMMENT 'author info',
     `source_site_id` bigint(20) unsigned NOT NULL COMMENT '来源网站',
     `language` varchar(20) NOT NULL DEFAULT '' COMMENT '文章语言类型',
     `publish_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'publish time',
     `url` varchar(1024) NOT NULL DEFAULT '' COMMENT '文章链接',
     `type` varchar(20) NOT NULL DEFAULT '' COMMENT '文章类别',
     `title` text NOT NULL COMMENT '标题',
     `content` longtext NOT NULL COMMENT '正文内容',
     `status` tinyint(4) NOT NULL COMMENT 'status，1:init, 2:reject, 3:pass',
     `score` tinyint(4) NOT NULL COMMENT '分数',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `publish_at` (`publish_at`),
     KEY `type` (`type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article info';

CREATE TABLE `article_metas` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `article_id` bigint(20) unsigned NOT NULL COMMENT 'article id',
     `meta_type` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta type',
     `meta_key` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta key',
     `meta_value` varchar(255) NOT NULL DEFAULT '' COMMENT 'meta value',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `article_id` (`article_id`),
     KEY `meta` (`meta_key`, `meta_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article meta';

CREATE TABLE `article_summarys` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `article_id` bigint(20) unsigned NOT NULL,
     `title` text NOT NULL COMMENT '标题',
     `summary` text NOT NULL COMMENT '摘要',
     `trading_proposal` tinyint(4) NOT NULL COMMENT '买卖建议',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='article summary';

CREATE TABLE `summary_contents` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `summary_id` bigint(20) unsigned NOT NULL,
     `original` text NOT NULL COMMENT '原文总结',
     `translation` text NOT NULL COMMENT '翻译总结',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `article_id` (`summary_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='summary contents';

CREATE TABLE `summary_outlines` (
     `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
     `summary_id` bigint(20) unsigned NOT NULL,
     `title` text NOT NULL COMMENT '标题',
     `content` text NOT NULL COMMENT '内容',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time', `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time', PRIMARY KEY (`id`),
     KEY `article_id` (`summary_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='summary outlines';

CREATE TABLE `labels` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='labels 标签列';

CREATE TABLE `label_refs` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `source_id` bigint(20) unsigned NOT NULL,
     `source_type` varchar(31) NOT NULL DEFAULT '',
     `label_id` bigint(20) unsigned NOT NULL,
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
     PRIMARY KEY (`id`),
     KEY `source` (`source_id`, `source_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='labels 标签关联表';

alter table sites add column `type` varchar(20) NOT NULL DEFAULT '' COMMENT '网站类型' after `tag`;
alter table sites add column `type_key` varchar(128) NOT NULL DEFAULT '' COMMENT '不同类型网站对应的key，如八抓鱼的taskID' after `type`;
alter table sites modify column `type_key` varchar(1024) NOT NULL DEFAULT '';

CREATE TABLE `symbols` (
     `id` bigint(20) unsigned NOT NULL COMMENT 'UUID',
     `symbol` varchar(127) NOT NULL DEFAULT '' COMMENT '股票代码标识',
     `company` varchar(512) NOT NULL DEFAULT '' COMMENT '公司名称',
     `source` varchar(63) NOT NULL DEFAULT '' COMMENT '股票代码来源',
     `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'create time',
     `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'update time',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='symbols 股票代码';
