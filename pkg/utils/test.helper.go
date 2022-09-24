package utils

import (
	"context"
	"test_go_project/configs"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/repository"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadTestConfig(env string) *configs.Config {
	viper.SetConfigFile("../../configs/config.env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		logger.Info.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		logger.Error.Fatalln("Load config failed:", viper.ConfigFileUsed(), "\n", err)
	}
	if env != "" {
		viper.SetConfigName("config." + env)
		if err := viper.MergeInConfig(); err == nil {
			logger.Info.Println("Override config file with:", viper.ConfigFileUsed())
		} else {
			logger.Error.Fatalln("Override config failed:", viper.ConfigFileUsed(), "\n", err)
		}
	}
	config := configs.Load()
	return config
}
func SetupTestClient(env string) *mongo.Client {
	client, _ := repository.Setup(context.Background(), LoadTestConfig(env))
	return client
}
func SetupTestDB(env string, name string) *mongo.Database {
	client := SetupTestClient(env)
	db := client.Database(name)
	return db
}
