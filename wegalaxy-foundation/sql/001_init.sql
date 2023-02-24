use wgx_foundation;

-- This is needed when using DBLockManager
create table if not exists app_lock (
     lock_key varchar(255) not null primary key,
     lock_value varchar(255) not null,
     expired_at_unix_milli bigint not null,
     created_at timestamp(3) not null default current_timestamp(3),
     updated_at timestamp(3) null on update current_timestamp(3)
);

create table if not exists `user` (
     `id` varchar(255) not null primary key,
     `status` int not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3)
);

create table if not exists `user_tag` (
     `id` binary(16) not null primary key,
     `name` varchar(255) not null,
	`status` int not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     unique index `idx_name` using hash(name)
);

-- userinfo 和 usertag 的映射
create table if not exists `user_tag_mapping` (
     `user_id` varchar(255) not null,
     `user_tag_id` binary(16) not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     primary key `uniq_user_tag` (`user_id`, `user_tag_id`)
);

-- user roles 只有固定几种，需要提前导入
create table if not exists `user_role` (
     `id` binary(16) not null primary key,
     `name` varchar(255) not null,
     `image` varchar(255) not null,
     `status` int not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     unique index `idx_name` using hash(name)
);

-- user 和 role 的映射
create table if not exists `user_role_mapping` (
     `user_id` varchar(255) not null,
     `role_id` binary(16) not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     primary key `uniq_user_role` (`user_id`, `role_id`)
);
 
create table if not exists `space` (
     `id` binary(16) not null primary key,
     `name` varchar(255) not null,
     `description` varchar(255) null,
     `entry_point_type` varchar(255) null comment 'url, etc',
     `entry_point_url` varchar(255) null,
     `image_url` varchar(255),
     `status` int not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3)
);

create table if not exists `space_tag_mapping` (
     `user_tag_id` binary(16) not null,
     `space_id` binary(16) not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     primary key `uniq_space_tag` (`user_tag_id`, `space_id`)
);

create table if not exists `scheduler_group` (
     `id` binary(16) not null primary key,
     `name` varchar(255) not null,
     `description` varchar(255) null,
     `status` int not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     unique index `idx_name` using hash(name)
);

create table if not exists `scheduler_job` (
     `id` binary(16) not null primary key,
     `scheduler_group_id` binary(16) not null,
     `name` varchar(255) not null,
     `description` varchar(255) null,
     `status` int not null,
     `scheduler_type` varchar(255) not null comment 'CronStyle, DateBased, RecurrenceRule',
     `scheduler_definition` varchar(1024) not null comment 'in json format, content is based on type',
     `action` varchar(1024) not null comment 'in json format',
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3),
     unique index `idx_name` using hash(name)
);

create table if not exists `qr_code` (
     `id` binary(16) not null primary key,
     `space_id` binary(16) null,
     `display_name` varchar(255) not null,
     `usage_type` int not null comment 'limited usage or not',
     `remaining_usage` int null comment 'for limited usage',
     `status` int not null,
     `valid_from_unix_sec` bigint not null default 0 comment '0 means no limit',
     `valid_to_unix_sec` bigint not null default 0 comment '0 means no limit',
     `parameter_to_handler` varchar(1024) not null comment 'in json format',
     `action` varchar(255) not null,
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3)
);

-- 数据插入qr_code demo
-- INSERT INTO qr_code(id, space_id, display_name, usage_type, remaining_usage, status, valid_from_unix_sec, valid_to_unix_sec, parameter_to_handler, action, created_at, updated_at) 
-- VALUES (UUID_TO_BIN('ae7fba72-4560-11ed-b30e-6c240879a7e4'), UUID_TO_BIN('c1ff8336-3e6b-11ed-bac5-04e46a65bc7f'),'前端联队', 1, 0, 1, 0, 0, '{"medalName":"前端联队","jumpToUrl":"?path=/wagalaxy&map=前端联队"}', '1024CheckIn', current_timestamp(3),null);

create table if not exists `report` (
     `id` binary(16) not null primary key,
     `app_version` varchar(255) not null,
     `platform` int not null,
     `user_id`  varchar(255) not null,
     `login_type` int not null,
     `event_time` bigint not null,
     `oper_id` varchar(255) not null,
     `ext1` varchar(255),
     `ext2` varchar(255),
     `ext3` varchar(255),
     `ext4` varchar(255),
     `ext5` varchar(255),
     `ext6` varchar(255),
     `ext7` varchar(255),
     `ext8` varchar(255),
     `ext9` varchar(255),
     `ext10` varchar(255),
     `created_at` timestamp(3) not null default current_timestamp(3),
     `updated_at` timestamp(3) null on update current_timestamp(3)
);
-- ext 是预留字段，具体含义请查阅数据上报表