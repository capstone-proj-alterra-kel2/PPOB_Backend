package utils

import (
	"log"

	"github.com/spf13/viper"
)

func GetEnv(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error when read config file")
	}
	return viper.GetString(key)
}
