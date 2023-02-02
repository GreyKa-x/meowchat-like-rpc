package model

import "C"
import (
	"context"
	"fmt"
	"github.com/xh-polaris/meowchat-like-rpc/errorx"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/stores/mon"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

const SEGLOCK_LENGTH = 100

var Lock = new(SegLock)

type SegLock [SEGLOCK_LENGTH]sync.Mutex

func (sl *SegLock) hash(key string) uint64 {
	return hash.Hash([]byte(key)) % uint64(len(sl))
}

func (sl *SegLock) lock(ctx context.Context, key string) error {
	id := sl.hash(key)
	sl[id].Lock()
	select {
	case <-ctx.Done():
		sl[id].Unlock()
		return errorx.ErrOutOfTime
	default:
		return nil
	}
}

func (sl *SegLock) trylock(key string) bool {
	id := sl.hash(key)
	return sl[id].TryLock()
}

func (sl *SegLock) unlock(key string) {
	id := sl.hash(key)
	sl[id].Unlock()
}

const ScoreCollectionName = "score"

var _ ScoreModel = (*customScoreModel)(nil)

type (
	// ScoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScoreModel.
	ScoreModel interface {
		scoreModel
		ListTopK(ctx context.Context, k int64, last int, t string) ([]*pb.ItemScore, error)
		InsertMany(ctx context.Context, scores []Score) error
	}
	customScoreModel struct {
		*redis.Redis
		*defaultScoreModel
	}
)

func (c customScoreModel) InsertMany(ctx context.Context, scores []Score) error {
	createAt := time.Now()
	updateAt := time.Now()
	data := make([]interface{}, len(scores))
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

func (c customScoreModel) ListTopK(ctx context.Context, k int64, last int, t string) ([]*pb.ItemScore, error) {
	if last > 1 {
		// 长期数据
		cacheKey := fmt.Sprintf("cache:score_%s_%d_%d", t, last, k)
		err := Lock.lock(ctx, cacheKey)
		if err != nil {
			return nil, err
		}
		defer Lock.unlock(cacheKey)
		d, err := c.Redis.ZrangeWithScoresCtx(ctx, cacheKey, 0, -1)
		if err != nil {
			return nil, err
		}
		if len(d) != 0 {
			// 有缓存数据
			return pairToItems(d), nil

		} else {
			// 无缓存数据
			data, err := c.ListTopKFromDb(ctx, k, last, t)
			if err != nil {
				return nil, err
			}
			_, err = c.Redis.ZaddsCtx(ctx, cacheKey, aggrToPair(data)...)
			if err != nil {
				return nil, err
			}
			err = c.Redis.Expire(cacheKey, getExpireTime(int64(last)))
			if err != nil {
				return nil, err
			}
			return toItems(data), nil
		}

	} else {
		// 当日数据
		data, err := c.Redis.ZrangeWithScoresCtx(ctx, "score_"+t, -k, -1)
		if err != nil {
			return nil, err
		}
		return pairToItems(data), nil
	}
}

func (c customScoreModel) ListTopKFromDb(ctx context.Context, k int64, last int, t string) ([]*AggrScore, error) {

	var data []*AggrScore
	wanted := time.Now().AddDate(0, 0, -last)
	matchStage := bson.D{{"$match", bson.D{{"type", t}, {"createAt", bson.M{"$gte": wanted}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$sid"}, {"total_score", bson.M{"$sum": "$score"}}}}}
	sortStage := bson.D{{"$sort", bson.M{"total_score": -1}}}
	limitStage := bson.D{{"$limit", k}}

	allowDiskUse := true
	opts := options.AggregateOptions{AllowDiskUse: &allowDiskUse}
	err := c.conn.Aggregate(ctx, &data, mongo.Pipeline{matchStage, groupStage, sortStage, limitStage}, &opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// NewScoreModel returns a model for the mongo.
func NewScoreModel(url, db, collection string, cache *redis.Redis) ScoreModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customScoreModel{
		defaultScoreModel: newDefaultScoreModel(conn),
		Redis:             cache,
	}
}
