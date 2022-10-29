package config

import (
	"github.com/spf13/viper"
	"log"
)

var Config ConfigStruct

func init() {
	var config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalln("Config Error")
	}
	err = config.Unmarshal(&Config)
	if err != nil {
		log.Fatalln("Config Error")
	}
}
