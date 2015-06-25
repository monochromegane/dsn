package dsn

import "testing"

func TestToDsnString(t *testing.T) {
	dsn := mysqlDsn{
		username: "username",
		password: "password",
		dbname:   "dbname",
		host:     "host",
		port:     "0",
		options:  options{"key": "value"},
	}

	expect := "username:password@tcp(host:0)/dbname?key=value"
	if actual := dsn.toDsnString(); actual != expect {
		t.Errorf("dsn should be %s, but %s", expect, actual)
	}
}

func TestToDsnString_Default(t *testing.T) {
	dsn := mysqlDsn{
		username: "username",
		password: "password",
		dbname:   "dbname",
	}

	expect := "username:password@tcp(localhost:3306)/dbname"
	if actual := dsn.toDsnString(); actual != expect {
		t.Errorf("dsn should be %s, but %s", expect, actual)
	}
}

func TestToDsnString_Socket(t *testing.T) {
	dsn := mysqlDsn{
		username: "username",
		password: "password",
		dbname:   "dbname",
		socket:   "test.sock",
	}

	expect := "username:password@unix(test.sock)/dbname"
	if actual := dsn.toDsnString(); actual != expect {
		t.Errorf("dsn should be %s, but %s", expect, actual)
	}
}
