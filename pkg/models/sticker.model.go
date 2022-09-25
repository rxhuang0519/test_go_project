package models

import (
	"encoding/json"
	"test_go_project/pkg/logger"
)

type Sticker struct {
	*Base       `json:",inline" bson:",inline"`
	StickerType string `json:"stickerType,omitempty" bson:"stickerType,omitempty"`
	StickerId   string `json:"stickerId,omitempty" bson:"stickerId,omitempty"`
	PackageId   string `json:"packageId,omitempty" bson:"packageId,omitempty"`
}

func NewStricker(stcType string, stcId string, pkgId string) *Sticker {
	return &Sticker{
		Base:        NewBase(),
		StickerType: stcType,
		StickerId:   stcId,
		PackageId:   pkgId,
	}
}

func (obj *Sticker) String() string {
	res, err := json.Marshal(obj)
	if err != nil {
		logger.Error.Println("JSON Error:", err)
	}
	return string(res)
}
