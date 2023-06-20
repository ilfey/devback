-- 
-- ПРИ ИЗМЕНЕНИИ ОГРАНИЧЕНИЙ НУЖНО МЕНЯТЬ,
-- ТАКЖЕ НУЖНО МЕНЯТЬ ОГРАНИЧЕНИЯ В ВАЛИДАЦИИ ДАННЫХ
-- СМ. ./iternal/pkg/models/** и ./internal/app/endpoints/v*/**
-- 


CREATE TABLE IF NOT EXISTS users(
  user_id varchar(16) primary key,
  password varchar(64) not null,
  is_deleted boolean not null default false,
  created_at timestamp default now(),
  modified_at timestamp default now()
);

CREATE TABLE IF NOT EXISTS messages(
  message_id bigint primary key generated always as identity, 
  content varchar(2000) not null,
  fk_user_id varchar(16) not null,
  fk_reply_message_id bigint,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false,
  foreign key(fk_reply_message_id) references messages(message_id) on delete set null,
  foreign key(fk_user_id) references users(user_id) on delete cascade
);

CREATE TABLE IF NOT EXISTS links(
  link_id bigint primary key generated always as identity,
  description varchar(256),
  url varchar(2000) not null,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false
);

CREATE TABLE IF NOT EXISTS contacts(
  contact_id bigint primary key generated always as identity,
  title varchar not null,
  fk_link_id bigint not null,
  modified_at timestamp default now(),
  created_at timestamp default now(),
  is_deleted boolean not null default false,
  foreign key(fk_link_id) references links(link_id) on delete cascade on update cascade
);