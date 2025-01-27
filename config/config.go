package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	configPath := filepath.Join(cwd, "../..", "config.yml")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist at: %s", configPath)
	}

	config.SetConfigFile(configPath)
	config.SetConfigType("yaml")

	err = config.ReadInConfig()

	log.Printf("Loaded config: %+v", config.AllSettings())

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return config
}
