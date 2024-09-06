CREATE TABLE IF NOT EXISTS users17 (
  id serial PRIMARY KEY,
  username varchar(50),
  email varchar(50) UNIQUE,
  password varchar(100)
);

CREATE TABLE IF NOT EXISTS tasks17 (
  id serial PRIMARY KEY,
  user_id int references users17(id) ON DELETE CASCADE,
  title varchar(50),
  created_at varchar(50),
  updated_at varchar(50) DEFAULT ''
);