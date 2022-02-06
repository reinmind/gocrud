package config

import (
	"log"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("server.port", "3000")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(`file not found: %v`, err)
	}

}
