package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`
	DbHost     string `yaml:"db_host"`
	DbPort     string `yaml:"db_port"`
	DbName     string `yaml:"db_name"`
	DbDriver   string `yaml:"db_driver"`
}

var cfg *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}

		if err := cleanenv.ReadConfig("config.yml", cfg); err != nil {
			log.Fatal(err)
		}

	})

	return cfg
}
