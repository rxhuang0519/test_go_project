package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
)

type Message struct {
	*Base   `json:",inline" bson:",inline"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

func NewMessage(message string) *Message {
	return &Message{
		Base:    NewBase(),
		Message: message,
	}
}
func (obj *Message) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
