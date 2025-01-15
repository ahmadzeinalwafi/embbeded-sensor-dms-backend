package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Print the current directory to debug
	log.Printf("Current working directory: %s", cwd)

	// Set the path to the config file relative to the root (adjust the number of "..")
	configPath := filepath.Join(cwd, "config.env")

	// Check if the config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file does not exist at: %s", configPath)
	}

	// Set the config file and load it
	config.SetConfigFile(configPath)
	config.AutomaticEnv()

	err = config.ReadInConfig()

	log.Printf("Loaded config: %+v", config.AllSettings())

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return config
}
