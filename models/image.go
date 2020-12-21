package models

import (
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
)

type Image struct {
	mgm.DefaultModel `bson:",inline"`
	Id string `json:"id" bson:"id"`
	Data []byte `json:"data" bson:"data"`
}

func CreateImage(base64Data []byte, id string) error {
	user := &Image{
		Data: base64Data,
		Id: id,
	}

	return mgm.Coll(user).Create(user)
}

func GetImage(id string) (image *Image, error error) {
	res := &Image{}
	coll := mgm.Coll(image)

	err := coll.First(bson.M{"id": id}, res)
	return res, err
}
