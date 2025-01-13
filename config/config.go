package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigFile("config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()
	err := config.ReadInConfig()

	log.Printf("Loaded config: %+v", config.AllSettings())

	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return config
}
