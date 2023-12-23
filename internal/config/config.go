package config

import "github.com/spf13/viper"

func Init() {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic("read env file fail")
	}
}
