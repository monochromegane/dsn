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
	dsn, err := adapter.dsnFromReader(r, "development")
	if err != nil {
		t.Errorf("error shouldn't occur, but %v", err)
	}

	if actual, ok := dsn.(mysqlDsn); !ok {
		t.Errorf("dsn type should mysql, but %v", actual)
	}
}
