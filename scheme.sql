-- 
-- ПРИ ИЗМЕНЕНИИ ОГРАНИЧЕНИЙ НУЖНО МЕНЯТЬ,
-- ТАКЖЕ НУЖНО МЕНЯТЬ ОГРАНИЧЕНИЯ В ВАЛИДАЦИИ ДАННЫХ
-- СМ. ./iternal/pkg/models/** и ./internal/app/endpoints/v*/**
-- 

CREATE TABLE IF NOT EXISTS users(
  username varchar(16) primary key,
  password varchar(64) not null,
  is_deleted boolean not null default false,
  created_at timestamp default now(),
  modified_at timestamp default now()
);

CREATE TABLE IF NOT EXISTS messages(
  id bigint primary key generated always as identity,
  content varchar(2000) not null,
  userId varchar(16) not null,
  reply bigint,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false,
  foreign key(reply) references messages(id) on delete set null,
  foreign key(userId) references users(username)
);

CREATE TABLE IF NOT EXISTS links(
  id bigint primary key generated always as identity,
  description varchar(256),
  url varchar(2000) not null,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false
);

CREATE TABLE IF NOT EXISTS contacts(
  id bigint primary key generated always as identity,
  title varchar not null,
  linkId bigint not null,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false,
  foreign key(linkId) references links(id) on delete cascade on update cascade
);