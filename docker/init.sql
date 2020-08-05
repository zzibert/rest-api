CREATE TABLE [IF NOT EXISTS] groups (
  id serial PRIMARY KEY,
  name varchar(50) UNIQUE NOT NULL,
);

CREATE TABLE [IF NOT EXISTS] users (
  id serial PRIMARY KEY,
  name varchar(50) UNIQUE NOT NULL,
  password varchar(50) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  group_id integer references groups(group_id)
);