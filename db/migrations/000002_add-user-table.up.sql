START TRANSACTION;

CREATE TABLE users (
  id uuid NOT NULL DEFAULT (uuid_generate_v4()) PRIMARY KEY,
  account_id uuid NOT NULL,
  full_name VARCHAR(255) NOT NULL,
  bio VARCHAR(255),
  created_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  updated_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  deleted_at pg_catalog.timestamptz
);

ALTER TABLE "users" ADD CONSTRAINT fk_users_accountId FOREIGN KEY (account_id) REFERENCES public."accounts"(id);

COMMIT;