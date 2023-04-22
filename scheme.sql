-- 
-- ПРИ ИЗМЕНЕНИИ ОГРАНИЧЕНИЙ НУЖНО МЕНЯТЬ,
-- ТАКЖЕ НУЖНО МЕНЯТЬ ОГРАНИЧЕНИЯ В ВАЛИДАЦИИ ДАННЫХ
-- СМ. ./iternal/pkg/models/*
-- 

CREATE TABLE IF NOT EXISTS users(
  username varchar(16) primary key,
  password varchar(64) not null,
  created_at timestamp default now()
);

CREATE TABLE IF NOT EXISTS messages(
  id bigint primary key generated always as identity,
  content varchar(2000) not null,
  userId varchar(16) not null,
  reply bigint,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  foreign key(reply) references messages(id),
  foreign key(userId) references users(username)
);
