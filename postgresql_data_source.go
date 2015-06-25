package dsn

import "strings"

type postgresqlDataSource struct {
	user     string
	password string
	host     string
	port     string
	dbname   string
}

func (d postgresqlDataSource) hostOrDefault() string {
	if d.host == "" {
		return "localhost"
	}
	return d.host
}

func (d postgresqlDataSource) portOrDefault() string {
	if d.port == "" {
		return "5432"
	}
	return d.port
}

func (d postgresqlDataSource) toDsnString() (string, string) {
	var values []string
	if d.dbname != "" {
		values = append(values, "dbname="+d.dbname)
	}
	if d.user != "" {
		values = append(values, "user="+d.user)
	}
	if d.password != "" {
		values = append(values, "password="+d.password)
	}
	values = append(values, "host="+d.hostOrDefault())
	values = append(values, "port="+d.portOrDefault())

	return "postgres", strings.Join(values, " ")
}
