package config

import (
	"fmt"
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
			DBHost:     "localhost",
			DBPort:     "3306",
			DBUserName: "root",
			DBPassword: "root",
			DBName:     "ca_mission",
		}
	})

	return appConfig
}
