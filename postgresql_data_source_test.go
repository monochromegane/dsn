package dsn

import "testing"

func TestToDsnString_PostgreSql(t *testing.T) {
	ds := postgresqlDataSource{
		user:     "user",
		password: "password",
		dbname:   "dbname",
		host:     "host",
		port:     "0",
	}

	expectName := "postgres"
	expectDsn := "dbname=dbname user=user password=password host=host port=0"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}

func TestToDsnString_PostgreSql_Default(t *testing.T) {
	ds := postgresqlDataSource{
		user:     "user",
		password: "password",
		dbname:   "dbname",
	}

	expectName := "postgres"
	expectDsn := "dbname=dbname user=user password=password host=localhost port=5432"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}
