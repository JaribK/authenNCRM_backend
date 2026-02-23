package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost string
	ServerPort string
	Secret     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func ReadInConfig() Config {
	return Config{
		ServerHost: viper.GetString("SERVER_HOST"),
		ServerPort: viper.GetString("SERVER_PORT"),
		Secret:     viper.GetString("SECRET"),
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetString("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}
}

func InitConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
