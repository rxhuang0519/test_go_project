package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Base struct {
	Id        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"-" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"-" bson:"updatedAt,omitempty"`
}

func NewBase() *Base {
	return &Base{
		Id:        primitive.NewObjectID(),
		CreatedAt: time.Now().Truncate(time.Millisecond).UTC(),
		UpdatedAt: time.Now().Truncate(time.Millisecond).UTC(),
	}
}
func (obj *Base) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
