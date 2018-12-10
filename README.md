# migration

SQL migration tool

```console
go run cmd/main.go \
    -db="postgres://postgres@localhost:5432/dbname?sslmode=disable" \
    -source=./fixtures \
    -migrate=up
```