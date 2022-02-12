--liquibase formatted sql

--changeset me:1
create table test1
(
    id   bigserial primary key,
    name text
);

--changeset me:2
insert into test1 (id, name)
values (1, 'name 1');
insert into test1 (id, name)
values (2, 'name 2');

--changeset me:3
create sequence seq_test;
