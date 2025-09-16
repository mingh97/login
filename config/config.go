package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig *Config

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Cors struct {
		AllowOrigins []string
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config failed, err:%v\n", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unmarshal config failed, err:%v\n", err)
	}

	initDB()
}
