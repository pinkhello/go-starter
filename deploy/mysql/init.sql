create database if not exists db_test default charset utf8mb4;

use db_test;

create table if not exists business_group
(
    id         int(11)      not null auto_increment comment 'ID',
    name       varchar(255) not null comment '商户名称',
    status     varchar(10)  not null comment '状态｜valid｜invalid',
    created_at datetime     not null comment '创建时间',
    updated_at datetime     not null comment '更新时间',
    primary key (id)
) engine = innodb comment '商户表';