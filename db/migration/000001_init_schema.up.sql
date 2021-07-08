CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "todos" (
    "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT DEFAULT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "todos" ("title");