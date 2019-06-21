# go-pghelpers

Golang helpers for Postgres. It uses [pkg/pghelpers](https://github.com/pkg/go-pghelpers) under the hoods.


## Installation

```sh
go get github.com/neighborly/go-pghelpers
```

## Usage

### PostgresConfig

PostgresConfig holds the information about db connection. For example

```go
var postgresConfig = pghelpers.PostgresConfig{
    Host:       myTestString,
	Port:       1,
	Username:   myTestString,
	Password:   myTestString,
	Database:   myTestString,
	SSLEnabled: false,
	Tracing:    testTracingConfig
}
```

### Connect Postgres

pghelpers.ConnectPostgres function returns db object and error message if any. You can use pghelpers.ConnectPostgres as follows

```go
db, err := pghelpers.ConnectPostgres(postgresConfig)
```

### GenerateAddress

GenerateAddress function returns Postgres connection string, which is used for sql.open().

```go
addr := postgresConfig.GenerateAddress()
```

## License

This project is licensed under the [MIT License](LICENSE.md).
