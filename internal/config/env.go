package config

import (
"os"
)

type Env struct {
	Password string
	Username string
	DBName   string
	Hostname string
	Mode     string
}

// LoadEnv returns all env variables needed in code
func LoadEnv() (env Env) {
	env.Password = os.Getenv("DB_PASSWORD")
	env.Username = os.Getenv("DB_USERNAME")
	env.DBName = os.Getenv("DB_NAME")
	env.Hostname = os.Getenv("DB_HOSTNAME")
	env.Mode = os.Getenv("DB_MODE")

	return
}
