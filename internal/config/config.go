package config

import "os"

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		DBname   string
	}
	Email struct {
		Sender string
		Password string
	}
}

func Configuration() *Config {
	c := &Config{}

	c.Database.User = osGetenv("DB_USER", "postgres")
	c.Database.Password = osGetenv("DB_PASSWORD", "2005")
	c.Database.Host = osGetenv("DB_HOST", "postgres")
	c.Database.Port = osGetenv("DB_PORT", "5432")
	c.Database.DBname = osGetenv("DB_NAME", "internship")
	c.Email.Sender = osGetenv("EMAIL_SENDER", "dostonxoshimov2005@gmail.com")
	c.Email.Password = osGetenv("EMAIL_PASSWORD", "jsxd uzpp wttr pwvk")

	return c
}

func osGetenv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
