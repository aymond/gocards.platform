DROP TABLE IF EXISTS  games;
DROP TABLE IF EXISTS  sessions;
DROP TABLE IF EXISTS  users;


create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null   
);

create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);

create table games (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  game_name  varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null   
);