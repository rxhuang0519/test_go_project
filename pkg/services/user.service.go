package services

import (
	"context"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Service[models.User]
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		Service: Service[models.User]{
			collection: db.Collection("users"),
		},
	}
}
func (service *UserService) FindByUserId(ctx context.Context, userId string) ([]*models.User, error) {
	logger.Debug.Printf("[FindByUserId] Start. ( userId: %s )\n", userId)
	filter := bson.M{"userId": userId}
	result, err := service.Find(ctx, filter)
	logger.Debug.Println("[FindByUserId] Complete.")
	return result, err
}
func (service *UserService) FindByName(ctx context.Context, name string) ([]*models.User, error) {
	logger.Debug.Printf("[FindByUserId] Start. ( name: %s )\n", name)
	filter := bson.M{"name": name}
	opts := options.Find().SetSort(bson.D{{Key: "name", Value: 1}})
	result, err := service.Find(ctx, filter, opts)
	logger.Debug.Println("[FindByUserId] Complete.")
	return result, err
}
