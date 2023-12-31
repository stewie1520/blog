START TRANSACTION;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE accounts (
  id uuid NOT NULL DEFAULT (uuid_generate_v4()),
  email VARCHAR(255) NOT NULL,
  password VARCHAR(1000) NOT NULL,
  created_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  updated_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  deleted_at pg_catalog.timestamptz,
  CONSTRAINT pk_accounts PRIMARY KEY (id)
);

CREATE UNIQUE INDEX ix_accounts_email ON "public".accounts (email);

COMMIT;
