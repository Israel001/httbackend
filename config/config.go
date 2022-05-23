package config

import "os"

type DatabaseConfig struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUsername string
	DbPassword string
}

type Configuration struct {
	AppPort  string
	AppEnv   string
	Database *DatabaseConfig
}

func getEnvValue(original string, alternative string) string {
	if original == "" {
		return alternative
	}
	return original
}

func New() *Configuration {
	c := new(Configuration)

	c.AppPort = getEnvValue(os.Getenv("APP_PORT"), "5000")
	c.AppEnv = getEnvValue(os.Getenv("APP_ENV"), "development")

	c.Database = new(DatabaseConfig)
	c.Database.DbHost = getEnvValue(os.Getenv("DB_HOST"), "127.0.0.1")
	c.Database.DbPort = getEnvValue(os.Getenv("DB_PORT"), "3306")
	c.Database.DbName = getEnvValue(os.Getenv("DB_DATABASE"), "htt")
	c.Database.DbUsername = getEnvValue(os.Getenv("DB_USERNAME"), "username")
	c.Database.DbPassword = getEnvValue(os.Getenv("DB_PASSWORD"), "password")

	return c
}
