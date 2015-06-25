package dsn

import "testing"

func TestToDsnString_Sqlite(t *testing.T) {
	ds := sqliteDataSource{
		dbname: "test.db",
	}

	expectName := "sqlite3"
	expectDsn := "test.db"
	name, dsn := ds.toDsnString()
	if name != expectName {
		t.Errorf("name should be %s, but %s", expectName, name)
	}
	if dsn != expectDsn {
		t.Errorf("dsn should be %s, but %s", expectDsn, dsn)
	}
}
