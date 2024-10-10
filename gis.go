package gis

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GisSmt5 struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Province    string             `json:"province,omitempty" bson:"province,omitempty"`
	District    string             `json:"district,omitempty" bson:"district,omitempty"`
	SubDistrict string             `json:"sub_district,omitempty" bson:"sub_district,omitempty"`
	Village     string             `json:"village,omitempty" bson:"village,omitempty"`
	Border      [][]float64        `json:"border,omitempty" bson:"border,omitempty"`
}
