package utils

import (
	"test_go_project/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToDoc(inputs interface{}) interface{} {
	data, err := bson.Marshal(inputs)
	if err != nil {
		logger.Error.Println("BSON Marshal Error:", err)
	}
	logger.Debug.Println("data:", data)
	var docs interface{}
	err = bson.Unmarshal(data, &docs)
	if err != nil {
		logger.Error.Println("BSON Unmarshal Error:", err)
	}
	logger.Debug.Println("docs:", docs)
	return &docs
}
func ToOId(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error.Println("Invalid ObjectId:", id)
	}
	return oid
}
