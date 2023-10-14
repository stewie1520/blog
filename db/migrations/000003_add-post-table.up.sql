START TRANSACTION;

CREATE TABLE posts (
  id uuid NOT NULL DEFAULT (uuid_generate_v4()),
  content TEXT NOT NULL,
  user_id uuid NOT NULL,

  created_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  updated_at pg_catalog.timestamptz NOT NULL DEFAULT (now()),
  deleted_at pg_catalog.timestamptz,
  CONSTRAINT pk_posts PRIMARY KEY (id),
  CONSTRAINT fk_posts_user_Id FOREIGN KEY (user_id) REFERENCES public."users"(id)
);

COMMIT;