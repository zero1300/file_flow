package config

import (
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()

	if err != nil {
		panic("load config error")
	}

}
