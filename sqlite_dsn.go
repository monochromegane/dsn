package dsn

type sqliteDsn struct {
	username string
	password string
	host     string
	port     string
	socket   string
	dbname   string
	options  options
}
