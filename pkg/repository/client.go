package repository

import (
	"context"
	"fmt"
	"test_go_project/configs"
	"test_go_project/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const CONNECT_TIME_OUT = 5 * time.Second

func generateConnectOptions(config *configs.Config) string {
	options := ""
	return options
}
func generateURI(config *configs.Config) string {
	options := generateConnectOptions(config)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, options)
	logger.Debug.Println("uri:", uri)
	return uri
}
func NewClient(ctx context.Context, config *configs.Config) *mongo.Client {
	logger.Info.Println("Setup DB...")
	dbUri := generateURI(config)
	client, err := mongo.NewClient(options.Client().ApplyURI(dbUri))
	ctx, cancel := context.WithTimeout(ctx, CONNECT_TIME_OUT)
	defer cancel()
	if err != nil {
		logger.Error.Panicln("Setup DB client failed:\n", err)
	}
	Connect(ctx, client)
	logger.Info.Println("Setup DB complete.")
	return client
}

func Connect(ctx context.Context, client *mongo.Client) {
	logger.Info.Println("Connecting To DB...")
	if err := client.Connect(ctx); err != nil {
		logger.Error.Panicln("Connection failed:\n", err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Error.Panicln("Connect failed:\n", err)
	}

	logger.Info.Println("Connection complete.")

}
func Disconnect(client *mongo.Client) error {
	logger.Info.Println("Disconnecting DB...")
	err := client.Disconnect(context.TODO())
	if err != nil {
		logger.Error.Panicln("Disconnect failed:\n", err)
	}
	logger.Info.Println("Disconnect complete.")
	return nil
}
