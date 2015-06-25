package dsn

import (
	"fmt"
	"strings"
)

type mysqlDataSource struct {
	username string
	password string
	host     string
	port     string
	socket   string
	dbname   string
	options  options
}

type options map[string]string

func (o options) toString() string {
	if len(o) == 0 {
		return ""
	}
	var values []string
	for k, v := range o {
		values = append(values, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(values, "&")
}

func (d mysqlDataSource) address() string {
	switch {
	case d.socket != "":
		return fmt.Sprintf("unix(%s)", d.socket)
	default:
		return fmt.Sprintf("tcp(%s:%s)", d.hostOrDefault(), d.portOrDefault())
	}
}

func (d mysqlDataSource) hostOrDefault() string {
	if d.host == "" {
		return "localhost"
	}
	return d.host
}

func (d mysqlDataSource) portOrDefault() string {
	if d.port == "" {
		return "3306"
	}
	return d.port
}

func (d mysqlDataSource) toDsnString() (string, string) {
	return "mysql",
		fmt.Sprintf("%s:%s@%s/%s%s",
			d.username,
			d.password,
			d.address(),
			d.dbname,
			d.options.toString(),
		)
}
