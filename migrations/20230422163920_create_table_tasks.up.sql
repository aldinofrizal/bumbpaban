CREATE TABLE IF NOT EXISTS tasks (
  id serial PRIMARY KEY,
  title varchar NOT NULL,
  description varchar NOT NULL,
  board_id int NOT NULL,
  status int NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  FOREIGN KEY ("board_id")
    REFERENCES boards ("id")
)