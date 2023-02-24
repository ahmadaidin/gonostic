package config

import (
	"log"
	"sync"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	cfg     configuration
	cfgOnce sync.Once
	envFile *string
)

type configuration struct {
	Environment               string        `env:"ENV" env-default:"prod" env-upd:"true"`
	Port                      int           `env:"PORT" env-default:"8000"`
	DatabaseURI               string        `env:"DATABASE_URI" env-upd:"true" env-required:"true"`
	DatabaseConnectionTimeout time.Duration `env:"DATABASE_CONNECTION_TIMEOUT" env-default:"10s" env-upd:"true"`
}

// Read reads the configuration file and sets the envFile variable
// If the file is not found, it will try to read the file from enviroment variable
func Read(file string) {
	cfgOnce.Do(func() {
		envFile = &file
		log.Printf(`Reading config file: "%s"`, *envFile)
		err := cleanenv.ReadConfig(file, &cfg)
		if err != nil {
			file = ".env"
			err := cleanenv.ReadConfig(file, &cfg)
			if err != nil {
				err := cleanenv.ReadEnv(&cfg)
				if err != nil {
					log.Fatalf("Config error %s", err.Error())
				}
				fileFlag := "nofile"
				envFile = &fileFlag
			} else {
				envFile = &file
			}
		}
	})
}

func GetConfig() configuration {
	if envFile == nil {
		log.Panic(`configuration file is not set. Call ReadConfig("path_to_file") first. path_to_file could be an empty string`)
	}
	err := cleanenv.UpdateEnv(&cfg)
	if err != nil {
		log.Fatalf("Config error %s", err.Error())
	}
	return cfg
}
