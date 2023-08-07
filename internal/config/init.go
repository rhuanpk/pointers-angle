package config

import (
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal.Fatalln(err.Error())
	}
	viper.WatchConfig()

	if err := viper.UnmarshalKey("api", &API); err != nil {
		logger.Fatal.Fatalln(err.Error())
	}
	if err := viper.UnmarshalKey("db", &DB); err != nil {
		logger.Fatal.Fatalln(err.Error())
	}
}
