package config

import (
	"fmt"
	"os"
)

type Conf struct {
	DbUsername string
	DbPassword string
	DbIP       string
	DbPort     string
	DbDatabase string
}

func ReadConfigFromEnv() *Conf {
	return &Conf{
		DbUsername: getEnvOrDefault("DB_USERNAME", "root"),
		DbPassword: getEnvOrDefault("DB_PASSWORD", "test"),
		DbIP:       getEnvOrDefault("DB_IP", "localhost"),
		DbPort:     getEnvOrDefault("DB_PORT", "3306"),
		DbDatabase: getEnvOrDefault("DB_DATABASE", "backend"),
	}
}

func (c *Conf) GetFullDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.DbUsername,
		c.DbPassword,
		c.DbIP,
		c.DbPort,
		c.DbDatabase)
}

func getEnvOrDefault(key, def string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return def
}
