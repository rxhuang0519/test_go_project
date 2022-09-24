package services

import (
	"context"
	"test_go_project/pkg/logger"
	"test_go_project/pkg/models"
	"test_go_project/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IService[T models.Model] interface {
	Create(ctx context.Context, input *T) (*T, error)
	FindById(ctx context.Context, id string) (*T, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error)
	FindAll(ctx context.Context, opts ...*options.FindOptions) ([]*T, error)
	UpdateById(ctx context.Context, id string, input *T, opts ...*options.FindOneAndUpdateOptions)
	DeleteById(ctx context.Context, id string)
}
type Service[T models.Model] struct {
	collection *mongo.Collection
}

func (service *Service[T]) Create(ctx context.Context, input *T) (*T, error) {
	logger.Debug.Printf("[Create] Start. ( input: %+v )\n", input)
	idResult, err := service.collection.InsertOne(ctx, *input)
	if err != nil {
		logger.Error.Panicln("[Create] Insert Error:", err)
	}
	result, _ := service.FindOne(ctx, bson.M{"_id": idResult.InsertedID})
	logger.Debug.Println("[Create] Complete.")
	return result, err
}

func (service *Service[T]) FindById(ctx context.Context, id string) (*T, error) {
	logger.Debug.Printf("[FindById] Start. ( id: %s )\n", id)
	result, err := service.FindOne(ctx, bson.M{"_id": utils.ToOId(id)})
	logger.Debug.Println("[FindById] Complete.")
	return result, err
}

func (service *Service[T]) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	logger.Debug.Printf("[FindOne] Start. ( filter: %+v )\n", filter)
	var result *T
	err := service.collection.FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Debug.Println("[FindOne] Not Found.")
			return nil, nil
		}
		logger.Error.Panicln("[FindOne] Error:", err)
		return nil, err
	}
	logger.Debug.Println("[FindOne] Complete.")
	return result, nil
}
func (service *Service[T]) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*T, error) {
	logger.Debug.Printf("[Find] Start. ( filter: %+v, opts: %+v )\n", filter, opts)
	var result []*T
	cur, err := service.collection.Find(ctx, filter, opts...)
	if err != nil {
		logger.Error.Panicln("[Find] Find Error:", err)
		return nil, err
	}
	err = cur.All(ctx, &result)
	if err != nil {
		logger.Error.Panicln("[Find] Cursor Error:", err)
		return nil, err
	}
	logger.Debug.Println("[Find] Complete.")
	return result, nil
}
func (service *Service[T]) FindAll(ctx context.Context, opts ...*options.FindOptions) ([]*T, error) {
	logger.Debug.Printf("[FindAll] Start. ( opts: %+v )\n", opts)
	result, err := service.Find(ctx, bson.M{})
	logger.Debug.Println("[FindAll] Complete.")
	return result, err
}

func (service *Service[T]) UpdateById(ctx context.Context, id string, input *T, opts ...*options.FindOneAndUpdateOptions) (*T, error) {
	logger.Debug.Printf("[UpdateById] Start. ( id: %s, input: %+v, opts: %+v )\n", id, input, opts)
	filter := bson.M{"_id": utils.ToOId(id)}
	updateInput := bson.D{{Key: "$set", Value: input}}
	opts = append(opts, options.FindOneAndUpdate().SetReturnDocument(1))
	var result *T
	err := service.collection.FindOneAndUpdate(ctx, filter, updateInput, opts...).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		logger.Error.Panicln("[UpdateById] Error:", err)
		return nil, err
	}
	logger.Debug.Println("[UpdateById] Complete.")
	return result, nil
}
func (service *Service[T]) DeleteById(ctx context.Context, id string) error {
	logger.Debug.Printf("[DeleteById] Start. ( id: %s )\n", id)
	filter := bson.M{"_id": utils.ToOId(id)}
	delResult, err := service.collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error.Panicln("[DeleteById] Error:", err)
		return err
	}
	logger.Debug.Printf("[DeleteById] Complete. ( deletedCount = %v )\n", delResult.DeletedCount)
	return nil
}
