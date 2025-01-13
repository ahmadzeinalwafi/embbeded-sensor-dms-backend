// config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	return config
}
