package services

import (
	"context"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	Service[models.User]
}

func (service *UserService) setIndex() {
	service.collection.Indexes().CreateOne(context.TODO(), mongo.IndexModel{Keys: bson.D{{Key: "userId", Value: 1}}, Options: options.Index().SetUnique(true)})
}
func NewUserService(db *mongo.Database) *UserService {
	service := &UserService{
		Service: Service[models.User]{
			collection: db.Collection("users"),
		},
	}
	service.setIndex()
	return service
}
func (service *UserService) Upsert(ctx context.Context, userId string, input *models.User) (*models.User, error) {
	logger.Debug.Printf("[Upsert] Start. ( userId: %s )\n", userId)
	filter := bson.M{"userId": userId}
	input.Id = primitive.NilObjectID
	updateInput := bson.D{{Key: "$set", Value: input}}
	var result *models.User
	err := service.collection.FindOneAndUpdate(ctx, filter, updateInput, options.FindOneAndUpdate().SetReturnDocument(1).SetUpsert(true)).Decode(&result)
	if err != nil {
		logger.Error.Panicln("[Upsert] Error:", err)
		return nil, err
	}
	logger.Debug.Println("[Upsert] Complete.")
	return result, err
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
