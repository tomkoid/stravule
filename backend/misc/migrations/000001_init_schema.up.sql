CREATE TABLE IF NOT EXISTS users (
  -- id BIGSERIAL PRIMARY KEY,
  id serial NOT NULL PRIMARY KEY,
  user_hash text NOT NULL,
  sid text NOT NULL
);

-- CREATE INDEX IF NOT EXISTS "idx_users_id" ON "users" ("id");
