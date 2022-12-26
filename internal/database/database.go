package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DbConfig struct {
	AppPort  string `yaml:"APP_PORT"`
	Address  string `yaml:"ADDRESS"`
	User     string `yaml:"USER"`
	Name     string `yaml:"NAME"`
	Password string `yaml:"PASSWORD"`
}

func Connection(config DbConfig) *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.User, config.Password, config.Name))
	if err != nil {
		log.Fatal(err)
	}

	return db
}
