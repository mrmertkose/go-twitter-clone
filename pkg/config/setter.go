package config

import (
	"github.com/spf13/viper"
	"log"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}
}
