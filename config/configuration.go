package config

import (
	"os"
)

const (
	LOCAL      = "local"
	PRODUCTION = "production"
)

// ENVIRONMENT:
const ENVIRONMENT string = LOCAL // LOCAL, DEVELOPMENT, PRODUCTION //

var env = map[string]map[string]string{
	// local environment configuration
	"local": {
		"PIPELINE":     LOCAL,
		"PORT":         "8888",
		"MYSQL_HOST":   "localhost",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "root",
		"MYSQL_SCHEMA": "sbs",
	},
	"production": {
		"PIPELINE":     PRODUCTION,
		"PORT":         "8888",
		"MYSQL_HOST":   "mysql-db",
		"MYSQL_PORT":   "3306",
		"MYSQL_USER":   "root",
		"MYSQL_PASS":   "admin",
		"MYSQL_SCHEMA": "sbs",
	},
}

// CONFIG : global configuration
var CONFIG = env[ENVIRONMENT]

// Getenv : function for Environment Lookup
func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
