CREATE TABLE IF NOT EXISTS members(
  id serial PRIMARY KEY,
  user_id int NOT NULL,
  board_id int NOT NULL,
  role int NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP,
  FOREIGN KEY ("user_id")
    REFERENCES "users" ("id"),
  FOREIGN KEY ("board_id")
    REFERENCES boards ("id")
)