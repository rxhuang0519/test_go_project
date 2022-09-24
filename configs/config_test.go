package configs_test

import (
	"test_go_project/configs"
	"testing"

	"github.com/spf13/viper"
)

func TestLoad(t *testing.T) {
	cfgFile := "./config.env"
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		t.Log("Using config file:", viper.ConfigFileUsed())
	} else {
		t.Error("Load config failed:", viper.ConfigFileUsed(), "\n", err)
	}
	config := configs.Load()
	t.Logf("config result: %+v", config)
	t.Log("Passed.")
}
