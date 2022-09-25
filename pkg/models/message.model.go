package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
)

type Message struct {
	*Base     `json:",inline" bson:",inline"`
	Type      string   `json:"type,omitempty" bson:"type,omitempty"`
	UserId    string   `json:"userId,omitempty" bson:"userId,omitempty"`
	GroupId   string   `json:"groupId,omitempty" bson:"groupId,omitempty"`
	RoomId    string   `json:"roomId,omitempty" bson:"roomId,omitempty"`
	MessageId string   `json:"messageId,omitempty" bson:"messageId,omitempty"`
	Text      string   `json:"text,omitempty" bson:"text,omitempty"`
	Image     string   `json:"image,omitempty" bson:"image,omitempty"`
	Video     string   `json:"video,omitempty" bson:"video,omitempty"`
	Audio     string   `json:"audio,omitempty" bson:"audio,omitempty"`
	File      string   `json:"file,omitempty" bson:"file,omitempty"`
	Sticker   *Sticker `json:"sticker,omitempty" bson:"sticker,omitempty"`
}

func NewMessage(msgId string) *Message {
	return &Message{
		Base:      NewBase(),
		MessageId: msgId,
	}
}

func (obj *Message) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
