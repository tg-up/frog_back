package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerHost    string `mapstructure:"SERVER_HOST"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	Mode          string `mapstructure:"MODE"`
	AllowedOrigin string `mapstructure:"ALLOWED_ORIGIN"`
}

var GlobalConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	if err := viper.Unmarshal(&GlobalConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

}
