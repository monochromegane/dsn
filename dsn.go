package dsn

type Dsn struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func FromRailsConfig(file, env string) (string, string, error) {
	adapter := railsAdapter{}
	ds, err := adapter.dataSource(file, env)
	if err != nil {
		return "", "", err
	}
	name, dsn := ds.toDsnString()
	return name, dsn, nil
}
