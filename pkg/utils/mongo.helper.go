package utils

import (
	"encoding/json"
	"test_go_project/pkg/logger"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToObject(inputs interface{}) interface{} {
	data, err := json.Marshal(inputs)
	if err != nil {
		logger.Error.Println("JSON Marshal Error:", err)
	}
	// logger.Debug.Println("data:", data)
	var obj interface{}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		logger.Error.Println("JSON Unmarshal Error:", err)
	}
	// logger.Debug.Println("obj:", obj)
	return obj
}
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
