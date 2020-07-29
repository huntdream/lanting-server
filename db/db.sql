create database if not exists lanting;

use lanting;

create table if not exists articles
(
    id         int auto_increment primary key,
    title      varchar(128)       not null,
    excerpt    text,
    content    text               not null,
    author_id  int      default 1 null,
    deleted_at datetime           null,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime default null on update CURRENT_TIMESTAMP
);

create table if not exists users
(
    id       int auto_increment primary key,
    username varchar(32) unique not null,
    password varchar(256)       not null,
    avatar   varchar(512),
    email    varchar(512) unique
);

create table if not exists roles
(
    id          int auto_increment primary key,
    name        varchar(32) unique not null,
    description varchar(128)
);

insert into roles (id, name, description)
values ( 1,  "admin",  "Admin"),
       ( 2,  "editor",  "Editor"),
       ( 3,  "member",  "Member");


create table if not exists user_roles
(
    user_id int not null,
    role_id int not null
);