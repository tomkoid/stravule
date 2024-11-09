CREATE TABLE IF NOT EXISTS filters (
  id serial NOT NULL PRIMARY KEY,
  user_id integer NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  filter_text text NOT NULL,
  category text CHECK (category IN ('include', 'exclude')),
  created_at timestamp DEFAULT CURRENT_TIMESTAMP
);
