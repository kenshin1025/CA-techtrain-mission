package config

import (
	"fmt"
	"os"
	"sync"
)

var (
	appConfig *config
	once      sync.Once
)

type config struct {
	DBHost     string
	DBPort     string
	DBUserName string
	DBPassword string
	DBName     string
}

func (c *config) GenerateDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", c.DBUserName, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
}

func Config() *config {
	once.Do(func() {
		appConfig = &config{
			DBHost:     envOrDefault("CA_MISSION_DBHOST", "localhost"),
			DBPort:     envOrDefault("CA_MISSION_DBPORT", "3306"),
			DBUserName: envOrDefault("CA_MISSION_DBUSERNAME", "root"),
			DBPassword: envOrDefault("CA_MISSION_DBPASSWORD", "root"),
			DBName:     envOrDefault("CA_MISSION_DBNAME", "ca_mission"),
		}
	})

	return appConfig
}

func envOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
