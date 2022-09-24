package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
)

type User struct {
	*Base  `bson:",inline"`
	UserId string `bson:"userId,omitempty"`
	Name   string `bson:"name,omitempty"`
}

func NewUser(userId string) *User {
	return &User{
		Base:   NewBase(),
		UserId: userId,
	}
}
func (obj *User) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
