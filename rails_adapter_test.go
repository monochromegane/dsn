package dsn

import (
	"strings"
	"testing"
)

func TestDsnFromReader_RailsAdapter_MySql(t *testing.T) {
	r := strings.NewReader(`
development:
  adapter: mysql2
  encoding: utf8
  reconnect: false
  database: sample_development
  pool: 5
  username: username
  password: password
  host: localhost`)

	adapter := railsAdapter{}
	dsn, err := adapter.dataSourceFromReader(r, "development")
	if err != nil {
		t.Errorf("error shouldn't occur, but %v", err)
	}

	if actual, ok := dsn.(mysqlDataSource); !ok {
		t.Errorf("data source type should mysql, but %v", actual)
	}
}

func TestDsnFromReader_RailsAdapter_PostgreSql(t *testing.T) {
	r := strings.NewReader(`
development:
  adapter: postgresql
  encoding: utf8
  reconnect: false
  database: sample_development
  pool: 5
  username: username
  password: password
  host: localhost`)

	adapter := railsAdapter{}
	dsn, err := adapter.dataSourceFromReader(r, "development")
	if err != nil {
		t.Errorf("error shouldn't occur, but %v", err)
	}

	if actual, ok := dsn.(postgresqlDataSource); !ok {
		t.Errorf("data source type should postgresql, but %v", actual)
	}
}

func TestDsnFromReader_RailsAdapter_Sqlite(t *testing.T) {
	r := strings.NewReader(`
development:
  adapter: sqlite3
  database: test.db`)

	adapter := railsAdapter{}
	dsn, err := adapter.dataSourceFromReader(r, "development")
	if err != nil {
		t.Errorf("error shouldn't occur, but %v", err)
	}

	if actual, ok := dsn.(sqliteDataSource); !ok {
		t.Errorf("data source type should sqlite, but %v", actual)
	}
}
