package dsn

type Dsn struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

func FromRailsConfig(file, env string) (string, error) {
	adapter := railsAdapter{}
	dsn, err := adapter.dsn(file, env)
	if err != nil {
		return "", err
	}
	return dsn.toDsnString(), nil
}
