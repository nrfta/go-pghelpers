# go-pghelpers [![Build Status](https://travis-ci.com/neighborly/go-pghelpers.svg?branch=master)](https://travis-ci.com/neighborly/go-pghelpers)

Golang helpers for Postgres.

## Installation

```sh
go get github.com/neighborly/go-pghelpers
```

## Usage

### PostgresTracingConfig

`PostgresTracingConfig` holds configuration information for enabling tracing on database calls.

```go
var tracingConfig = pghelpers.PostgresTracingConfig{
	Enabled:			true,
	CreateRowsNextSpan:		true,
	CreateRowsCloseSpan:		true,
	CreateRowsAffectedSpan:		true,
	CreateLastInsertedIDSpan:	true,
	AddQueryAttribute:		true,
	AddQueryParamsAttributes:	true,
}
```

### PostgresConfig

`PostgresConfig` holds the information for postgres database connection. For example

```go
var postgresConfig = pghelpers.PostgresConfig{
	Host:       "localhost",
	Port:       5432,
	Username:   "postgres",
	Password:   "",
	Database:   "postgres",
	Tracing:    tracingConfig,
}
```

### Connect Postgres

`pghelpers.ConnectPostgres` function returns db object and error if any. You can use `pghelpers.ConnectPostgres` as follows

```go
db, err := pghelpers.ConnectPostgres(postgresConfig)
```

### GenerateAddress

`GenerateAddress` returns a Postgres connection string, which can be used for `sql.open()`.

```go
addr := postgresConfig.GenerateAddress()
```

## License

This project is licensed under the [MIT License](LICENSE.md).
