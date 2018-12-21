# migration
[![Build Status](https://travis-ci.org/gosidekick/migration.svg?branch=master)](https://travis-ci.org/gosidekick/migration)
[![Go Report Card](https://goreportcard.com/badge/github.com/gosidekick/migration)](https://goreportcard.com/report/github.com/gosidekick/migration)
[![GoDoc](https://godoc.org/github.com/gosidekick/migration?status.png)](https://godoc.org/github.com/gosidekick/migration)
[![Go project version](https://badge.fury.io/go/github.com%2Fgosidekick%2Fmigration.svg)](https://badge.fury.io/go/github.com/gosidekick/migration)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

SQL migration tool

```console
go run cmd/migration/main.go \
    -db="postgres://postgres@localhost:5432/dbname?sslmode=disable" \
    -source=./fixtures \
    -migrate=up
```

```console
go run cmd/migration/main.go -db="postgres://postgres@localhost:5432/dbname?sslmode=disable" -source=./fixtures -migrate=down
```



