create EXTENSION if not exists "uuid-ossp";

create table "todos" (
    "id" uuid primary key default uuid_generate_v4(),
    "title" varchar(255) not null,
    "description" text default null,
    "created_at" timestamptz not null default (now()),
    "updated_at" timestamptz not null default (now())
);

create index on "todos" ("title");