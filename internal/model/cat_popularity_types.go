package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CatPopularity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CatId      string             `bson:"catId,omitempty" json:"catId,omitempty"`
	Popularity int64              `bson:"popularity,omitempty" json:"popularity,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type CatPop struct {
	CatId      string `bson:"catId,omitempty" json:"catId,omitempty"`
	Popularity int64  `bson:"popularity,omitempty" json:"popularity,omitempty"`
}
