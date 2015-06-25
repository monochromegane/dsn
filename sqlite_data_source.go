package dsn

type sqliteDataSource struct {
	dbname string
}

func (d sqliteDataSource) toDsnString() (string, string) {
	return "sqlite3", d.dbname
}
