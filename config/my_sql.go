package config

type MySql struct {
	Connection string
}

func DbConnection() string {
	return cfg.MySql.Connection
}