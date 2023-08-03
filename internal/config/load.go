package config

import (
	"github.com/rhuanpk/pointers-angle/pkg/logger"

	"github.com/spf13/viper"
)

// Load load the config file "config.yml" in root directory.
func Load() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatal.Fatalln(err.Error())
	}

	if err := viper.Unmarshal(&Config); err != nil {
		logger.Fatal.Fatalln(err.Error())
	}
}
