package services

import (
	"context"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MessageService struct {
	Service[models.Message]
}

func NewMessageService(db *mongo.Database) *MessageService {
	return &MessageService{
		Service: Service[models.Message]{
			collection: db.Collection("messages"),
		},
	}
}
func (service *MessageService) FindByMessage(ctx context.Context, message string) ([]*models.Message, error) {
	logger.Debug.Printf("[FindByMessage] Start. ( message: %s )\n", message)
	filter := bson.M{"message": message}
	opts := options.Find().SetSort(bson.D{{Key: "message", Value: 1}})
	result, err := service.Find(ctx, filter, opts)
	logger.Debug.Println("[FindByMessage] Complete.")
	return result, err
}
