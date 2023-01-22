package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const ScoreCollectionName = "score"

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		ListTopK(ctx context.Context, k int64, last int, t string) ([]*AggrScore, error)
		InsertMany(ctx context.Context, scores []Score) error
	}
	customScoreModel struct {
		*defaultScoreModel
	}
)

func (c customScoreModel) InsertMany(ctx context.Context, scores []Score) error {
	createAt := time.Now()
	updateAt := time.Now()
	data := make([]interface{}, len(scores), len(scores))
	for i := 0; i < len(scores); i++ {
		if scores[i].ID.IsZero() {
			scores[i].ID = primitive.NewObjectID()
			scores[i].CreateAt = createAt
			scores[i].UpdateAt = updateAt
		}
		data[i] = scores[i]
	}
	_, err := c.conn.InsertMany(ctx, data)
	return err
}

func (c customScoreModel) ListTopK(ctx context.Context, k int64, last int, t string) ([]*AggrScore, error) {

	var data []*AggrScore
	wanted := time.Now().AddDate(0, 0, -last)
	matchStage := bson.D{{"$match", bson.D{{"type", t}, {"createAt", bson.M{"$gte": wanted}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$sid"}, {"total_score", bson.M{"$sum": "$score"}}}}}
	sortStage := bson.D{{"$sort", bson.M{"total_score": -1}}}
	limitStage := bson.D{{"$limit", k}}

	err := c.conn.Aggregate(ctx, &data, mongo.Pipeline{matchStage, groupStage, sortStage, limitStage})
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewScoreModel returns a model for the mongo.
func NewScoreModel(url, db, collection string) ScoreModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customScoreModel{
		defaultScoreModel: newDefaultScoreModel(conn),
	}
}
