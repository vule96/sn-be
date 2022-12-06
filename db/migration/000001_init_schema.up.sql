CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "posts" (
  "id" uuid default uuid_generate_v4(),
  "user_id" uuid not null,
  "content" varchar not null,
  "is_active" boolean not null default true,
  "created_at" timestamp with time zone not null default now(),
  "updated_at" timestamp with time zone default current_timestamp
);

CREATE INDEX ON "posts" ("user_id");