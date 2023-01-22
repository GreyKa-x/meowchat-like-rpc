package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Score struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type  string             `bson:"type,omitempty" json:"type,omitempty"`
	Sid   string             `bson:"sid,omitempty" json:"sid,omitempty"`
	Score int64              `bson:"score,omitempty" json:"score,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}

type AggrScore struct {
	Sid   string `bson:"_id,omitempty" `
	Score int64  `bson:"total_score"`
}
