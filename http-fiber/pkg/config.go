package pkg

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

var AppConfig Config

func InitConfig() {
	viper.SetConfigFile("config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Failed to read config file: %v, using default config", err)
		AppConfig = Config{}
		AppConfig.Server.Port = "3000"
	} else {
		err = viper.Unmarshal(&AppConfig)
		if err != nil {
			log.Fatalf("Failed to unmarshal config: %v", err)
		}
	}
}
