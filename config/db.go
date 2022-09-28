package config

import "fmt"

// database struct
// config struct(db struct)
//os package Getenv

// 2 func to intialise
type databaseConfig struct {
	driver       string
	host         string
	databasename string
	user         string
	password     string
	port         int
}

func Database() databaseConfig {
	return appConfig.db
}

func (c databaseConfig) Driver() string {
	return c.driver
}

func (c databaseConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", c.user, c.password, c.host, c.port, c.databasename)
}

func newDatabaseConfig() databaseConfig {
	return databaseConfig{
		driver:       readEnvString("DB_DRIVER"),
		host:         readEnvString("DB_HOST"),
		databasename: readEnvString("DB_NAME"),
		user:         readEnvString("DB_USER"),
		password:     readEnvString("DB_PASSWORD"),
		port:         readEnvInt("DB_PORT"),
	}
}
