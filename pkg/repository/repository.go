package repository

import (
	"context"
	"test_go_project/configs"
	"test_go_project/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewRepository(dbName string, config *configs.Config, client *mongo.Client) *Repository {
	if client == nil {
		client = NewClient(context.Background(), config)
	}
	return &Repository{
		client: client,
		db:     client.Database(dbName),
	}
}

func (service *Repository) SetDB(name string) *mongo.Database {
	service.db = service.client.Database(name)
	logger.Info.Println("Switch to DB:", service.db.Name())
	return service.db
}

func (service *Repository) DB() *mongo.Database {
	return service.db
}
