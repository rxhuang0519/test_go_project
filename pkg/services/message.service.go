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
func (service *MessageService) FindByText(ctx context.Context, text string) ([]*models.Message, error) {
	logger.Debug.Printf("[FindByText] Start. ( text: %s )\n", text)
	filter := bson.M{"text": text}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})
	result, err := service.Find(ctx, filter, opts)
	logger.Debug.Println("[FindByText] Complete.")
	return result, err
}
func (service *MessageService) FindByMessageId(ctx context.Context, messageId string) ([]*models.Message, error) {
	logger.Debug.Printf("[FindByMessageId] Start. ( messageId: %s )\n", messageId)
	filter := bson.M{"messageId": messageId}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})
	result, err := service.Find(ctx, filter, opts)
	logger.Debug.Println("[FindByMessageId] Complete.")
	return result, err
}
func (service *MessageService) FindByUserId(ctx context.Context, userId string) ([]*models.Message, error) {
	logger.Debug.Printf("[FindByUserId] Start. ( userId: %s )\n", userId)
	filter := bson.M{"userId": userId}
	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: 1}})
	result, err := service.Find(ctx, filter, opts)
	logger.Debug.Println("[FindByUserId] Complete.")
	return result, err
}
