package database

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type YAML struct {
	Configuration struct {
		Database struct {
			Hostname string `yaml:"host"`
			Username string `yaml:"user"`
			Password string `yaml:"password"`
			Database string `yaml:"dbname"`
			Port     string `yaml:"port"`
			SSLMode  string `yaml:"ssl_mode"`
			Timezone string `yaml:"timezone"`
		}
	}
}

func Connect(pathOfYaml string, gormConfig *gorm.Config) (database *gorm.DB, err error) {
	yamlBytes, err := os.ReadFile(pathOfYaml)
	if err != nil {
		return nil, err
	}

	configure := &YAML{}
	err = yaml.Unmarshal(yamlBytes, configure)
	if err != nil {
		return nil, err
	}

	formatString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s"
	dataSource := fmt.Sprintf(formatString, configure.Configuration.Database.Hostname, configure.Configuration.Database.Username, configure.Configuration.Database.Password, configure.Configuration.Database.Database, configure.Configuration.Database.Port, configure.Configuration.Database.SSLMode, configure.Configuration.Database.Timezone)
	database, err = gorm.Open(postgres.Open(dataSource), gormConfig)
	if err != nil {
		return nil, err
	}

	return database, nil
}
