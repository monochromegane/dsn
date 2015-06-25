package dsn

import (
	"bytes"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type railsAdapter struct{}

func (a railsAdapter) dataSource(file, env string) (dsnStringer, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return a.dataSourceFromReader(bytes.NewBuffer(data), env)
}

func (a railsAdapter) dataSourceFromReader(r io.Reader, env string) (dsnStringer, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var conf map[string]RailsDbConfig
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	switch conf[env].Adapter {
	case "mysql2":
		return mysqlDataSource{
			host:     conf[env].Host,
			port:     conf[env].Port,
			dbname:   conf[env].Database,
			username: conf[env].Username,
			password: conf[env].Password,
		}, nil
	case "postgresql":
		return postgresqlDataSource{
			host:     conf[env].Host,
			port:     conf[env].Port,
			dbname:   conf[env].Database,
			user:     conf[env].Username,
			password: conf[env].Password,
		}, nil
	case "sqlite3":
		return sqliteDataSource{
			dbname: conf[env].Database,
		}, nil
	}
	return nil, nil
}

type RailsDbConfig struct {
	Adapter  string
	Host     string
	Port     string
	Database string
	Username string
	Password string
}
