package config

import (
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetConfig(cfgFile string) *viper.Viper {
	config := viper.New()

	if cfgFile != "" {
		config.SetConfigFile(cfgFile)
	}
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	config.AutomaticEnv()

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		zap.S().Panicf("config file %s failed to load: %s.\n", cfgFile, err.Error())
	}

	return config
}
