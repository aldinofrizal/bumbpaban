# Bumpaban DOCS

## Migrations
- create migrations
```bash
migrate create -ext sql -dir migrations [MIGRATION_NAME]
```
- up migrations
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/bumpaban?sslmode=disable&TimeZone=Asia/Jakarta" -path migrations up
```

- down migrations
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/bumpaban?sslmode=disable&TimeZone=Asia/Jakarta" -path migrations dowm [STEP]
```

- force migrations version
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/bumpaban?sslmode=disable&TimeZone=Asia/Jakarta" -path migrations force [VERSION]
```