package config

import (
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"nokogiriwatir/radio-main/internal/database"
	"os"
)

type App struct {
	AppPort string `yaml:"APP_PORT"`
}

type Configs struct {
	Logger *zap.Config       `yaml:"logger"`
	Db     database.DbConfig `yaml:"db"`
	App    App               `yaml:"app"`
}

func InitConfig() Configs {
	fh, err := os.Open("app_dev.yaml")
	if err != nil {
		log.Fatal(err)
	}

	configs := Configs{}

	if er := yaml.NewDecoder(fh).Decode(&configs); er != nil {
		log.Fatalln(er)
	}
	zap.Must(configs.Logger.Build())
	return configs
}
