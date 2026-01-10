CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users (
  id uuid primary key default uuid_generate_v4()
  , created_at timestamptz default now()
  , name text
  , email text
  , passport text
);

create table IF NOT EXISTS task(
  id uuid primary key default uuid_generate_v4()
  , created_at timestamptz default now()
  , responsible_id uuid references users(id)
  , name text
  , status text
);