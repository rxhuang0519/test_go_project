package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreateAt time.Time          `json:"-" bson:"createdAt,omitempty"`
	UpdateAt time.Time          `json:"-" bson:"updatedAt,omitempty"`
}

func NewBase() *Base {
	return &Base{
		Id:       primitive.NewObjectID(),
		CreateAt: time.Now().Truncate(time.Millisecond).UTC(),
		UpdateAt: time.Now().Truncate(time.Millisecond).UTC(),
	}
}
func (obj *Base) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
