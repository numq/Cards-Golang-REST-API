package database

import (
	"fmt"
	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	AppPort    string `mapstructure:"APP_PORT"`
}

func (c Config) GetDbUrl() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.DbDriver, c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName)
}

func LoadConfig(path string) (config Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	AppConfig = config
	return AppConfig
}
