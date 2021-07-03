create database dev owner postgres;
\c dev;

drop table if exists users;
create table users (
    id serial primary key,
    name varchar(255)not null,
    created_at date not null default CURRENT_DATE,
    updated_at date not null default CURRENT_DATE
);

insert into users (name) values ('User1');