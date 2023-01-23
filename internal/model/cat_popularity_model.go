package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const PopularityCollectionName = "cat_pop"

var _ CatPopularityModel = (*customCatPopularityModel)(nil)

type (
	// CatPopularityModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCatPopularityModel.
	CatPopularityModel interface {
		catPopularityModel
		AddPopularity(ctx context.Context, catId string, val int64) error
		ListPopularity(ctx context.Context, catId []string) ([]CatPop, error)
		ListTopK(ctx context.Context, k int64) ([]CatPop, error)
	}

	customCatPopularityModel struct {
		*defaultCatPopularityModel
	}
)

func (c customCatPopularityModel) AddPopularity(ctx context.Context, catId string, val int64) error {
	filter := bson.D{{"catId", catId}}
	update := bson.D{{"$inc", bson.M{"popularity": val}}}

	upsert := true
	opts := options.UpdateOptions{Upsert: &upsert}
	_, err := c.conn.UpdateOne(ctx, filter, update, &opts)
	if err != nil {
		return err
	}
	return nil
}

func (c customCatPopularityModel) ListPopularity(ctx context.Context, catId []string) ([]CatPop, error) {
	filter := bson.D{{"catId", bson.M{"$in": catId}}}
	var data []CatPop
	err := c.conn.Find(ctx, &data, filter)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (c customCatPopularityModel) ListTopK(ctx context.Context, k int64) ([]CatPop, error) {

	sort := bson.D{{"popularity", -1}}
	opts := options.FindOptions{Sort: sort, Limit: &k}
	var data []CatPop
	err := c.conn.Find(ctx, &data, bson.D{}, &opts)
	if err != nil {
		return nil, err
	}
	return data, err
}

// NewCatPopularityModel returns a model for the mongo.
func NewCatPopularityModel(url, db, collection string) CatPopularityModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customCatPopularityModel{
		defaultCatPopularityModel: newDefaultCatPopularityModel(conn),
	}
}
