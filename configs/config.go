package configs

import (
	"test_go_project/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	DB_USER             string `mapstructure:"DB_USER"`
	DB_PASSWORD         string `mapstructure:"DB_PASSWORD"`
	DB_HOST             string `mapstructure:"DB_HOST"`
	DB_PORT             string `mapstructure:"DB_PORT"`
	LINE_CHANNEL_SECRET string `mapstructure:"LINE_CHANNEL_SECRET"`
	LINE_CHANNEL_TOKEN  string `mapstructure:"LINE_CHANNEL_TOKEN"`
	LINE_USER_ID        string `mapstructure:"LINE_USER_ID"`
}

func Load() *Config {
	var config *Config
	err := viper.Unmarshal(&config)
	if err != nil {
		logger.Error.Fatalln("Load config failed:\n", err)
	}
	logger.Info.Println("Load config complete.")
	return config
}
