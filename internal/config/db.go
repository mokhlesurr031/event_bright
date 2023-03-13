package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

var (
	db Database
)

func DB() *Database {
	return &db
}

func loadDB() {
	viper.SetConfigFile("config.yml")
	er := viper.ReadInConfig()
	if er != nil {
		log.Println(er)
	}

	currentDB := viper.GetString("current_db.RUNNING")
	if currentDB == "docker_database" {
		db = Database{
			Name:     viper.GetString(currentDB + ".POSTGRESQL_ADDON_DB"),
			Username: viper.GetString(currentDB + ".POSTGRESQL_ADDON_USER"),
			Password: viper.GetString(currentDB + ".POSTGRESQL_ADDON_PASSWORD"),
			Host:     viper.GetString(currentDB + ".POSTGRESQL_ADDON_HOST"),
			Port:     viper.GetInt(currentDB + ".POSTGRESQL_ADDON_PORT"),
		}

	} else if currentDB == "remote_db" {
		viper.AutomaticEnv()
		db = Database{
			Name:     viper.GetString("POSTGRESQL_ADDON_DB"),
			Username: viper.GetString(".POSTGRESQL_ADDON_USER"),
			Password: viper.GetString(".POSTGRESQL_ADDON_PASSWORD"),
			Host:     viper.GetString(".POSTGRESQL_ADDON_HOST"),
			Port:     viper.GetInt(".POSTGRESQL_ADDON_PORT"),
		}
		fmt.Println("DATABASEEEE", db)

	}
}
