# dsn [![Build Status](https://travis-ci.org/monochromegane/dsn.svg?branch=master)](https://travis-ci.org/monochromegane/dsn)

A data source name adapter for sql.Open in golang.

## Usage

```go
name, dsn, _ := dsn.FromRailsConfig("myapp/config/database.yml", "production")
fmt.Printf("[%s] %s", name, dsn)
// => [mysql] username:password@tcp(localhost:3306)/dbname
```

## Features

- Data Source Name from Rails database.yml.
- Support MySQL, PostgreSql, Sqlite.

## Installation

```sh
$ go get github.com/monochromegane/dsn
```

## TODO

- Data Source Name from flags.
- Data Source Name from Env.

## Contribution

1. Fork it
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

## License

[MIT](https://github.com/monochromegane/dsn/blob/master/LICENSE)

## Author

[monochromegane](https://github.com/monochromegane)

