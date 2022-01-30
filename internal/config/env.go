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
	env.Password = os.Getenv("PASSWORD")
	env.Username = os.Getenv("USERNAME")
	env.DBName = os.Getenv("DBNAME")
	env.Hostname = os.Getenv("HOSTNAME")
	env.Mode = os.Getenv("MODE")

	return
}
