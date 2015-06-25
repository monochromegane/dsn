package dsn

import "testing"

func TestToDsnString(t *testing.T) {
	ds := mysqlDataSource{
		username: "username",
		password: "password",
		dbname:   "dbname",
		host:     "host",
		port:     "0",
		options:  options{"key": "value"},
	}

	expectName := "mysql"
	expectDsn := "username:password@tcp(host:0)/dbname?key=value"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}

func TestToDsnString_Default(t *testing.T) {
	ds := mysqlDataSource{
		username: "username",
		password: "password",
		dbname:   "dbname",
	}

	expectName := "mysql"
	expectDsn := "username:password@tcp(localhost:3306)/dbname"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}

func TestToDsnString_Socket(t *testing.T) {
	ds := mysqlDataSource{
		username: "username",
		password: "password",
		dbname:   "dbname",
		socket:   "test.sock",
	}

	expectName := "mysql"
	expectDsn := "username:password@unix(test.sock)/dbname"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}
