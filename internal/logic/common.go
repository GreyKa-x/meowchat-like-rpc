package logic

import (
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
	"github.com/zeromicro/go-zero/core/hash"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math"
	"sync"
)

const SEGLOCK_LENGTH = 100

var once sync.Once

type SegLock struct {
	locks []sync.Mutex
}

var Lock *SegLock = new(SegLock)

func (sl *SegLock) hashval(key string) uint64 {
	once.Do(func() {
		sl.locks = make([]sync.Mutex, SEGLOCK_LENGTH)
	})
	return hash.Hash([]byte(key)) % uint64(len(sl.locks))
}

func (sl *SegLock) lock(key string) {
	id := sl.hashval(key)
	sl.locks[id].Lock()
}
func (sl *SegLock) trylock(key string) {
	id := sl.hashval(key)
	sl.locks[id].TryLock()
}

func (sl *SegLock) unlock(key string) {
	id := sl.hashval(key)
	sl.locks[id].Unlock()
}

func toCatPop(cats []model.CatPop) []*pb.CatPop {
	ret := make([]*pb.CatPop, len(cats))
	for i := 0; i < len(cats); i++ {
		ret[i] = &pb.CatPop{
			CatId:      cats[i].CatId,
			Popularity: cats[i].Popularity,
		}
	}
	return ret
}

func getExpireTime(days int64) int {
	d := math.Pow(float64(days), 3.0/4) / 7
	return int(d * 3600 * 24)
}

func aggrToPair(m []*model.AggrScore) []redis.Pair {
	r := make([]redis.Pair, len(m)+1)
	for i := 0; i < len(m); i++ {
		r[i].Key = m[i].Sid
		r[i].Score = m[i].Score
	}
	r[len(m)].Key = "_"
	r[len(m)].Score = math.MinInt64
	return r
}

func pairToItems(m []redis.Pair) []*pb.ItemScore {
	r := make([]*pb.ItemScore, 0, len(m))
	for i := 0; i < len(m); i++ {
		if m[i].Score == math.MinInt64 {
			continue
		}
		r = append(r, &pb.ItemScore{
			Id:    m[i].Key,
			Score: m[i].Score,
		})
	}
	return r
}

func toItems(m []*model.AggrScore) []*pb.ItemScore {
	r := make([]*pb.ItemScore, len(m))
	for i := 0; i < len(m); i++ {
		r[i] = &pb.ItemScore{
			Id:    m[i].Sid,
			Score: m[i].Score,
		}
	}
	return r
}
