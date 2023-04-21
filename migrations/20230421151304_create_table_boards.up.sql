CREATE TABLE IF NOT EXISTS boards (
  id serial PRIMARY KEY,
  title varchar NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
)