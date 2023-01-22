package logic

import (
	"context"
	"fmt"
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"math"

	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTopKLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTopKLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTopKLogic {
	return &GetTopKLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTopKLogic) GetTopK(in *pb.GetTopKReq) (*pb.GetTopKResp, error) {

	if in.Range > 1 {
		// 长期数据
		cachekey := fmt.Sprintf("cache:score_%s_%d_%d", in.Type, in.Range, in.K)
		d, err := l.svcCtx.Redis.ZrangeWithScoresCtx(l.ctx, cachekey, 0, -1)
		if err != nil {
			return nil, err
		}
		if len(d) != 0 {
			// 有缓存数据
			return &pb.GetTopKResp{Item: pairToItems(d)}, nil

		} else {
			// 无缓存数据
			data, err := l.svcCtx.ScoreModel.ListTopK(l.ctx, in.K, int(in.Range), in.Type)
			if err != nil {
				return nil, err
			}
			_, err = l.svcCtx.Redis.ZaddsCtx(l.ctx, cachekey, aggrToPair(data)...)
			err = l.svcCtx.Redis.Expire(cachekey, getExpireTime(in.Range))
			if err != nil {
				return nil, err
			}
			return &pb.GetTopKResp{Item: toItems(data)}, nil
		}

	} else {
		// 当日数据
		data, err := l.svcCtx.Redis.ZrangeWithScoresCtx(l.ctx, "score_"+in.Type, -in.K, -1)
		if err != nil {
			return nil, err
		}
		return &pb.GetTopKResp{Item: pairToItems(data)}, nil
	}

}

func getExpireTime(days int64) int {
	d := math.Pow(float64(days), 3/4) / 7
	return int(d * 3600 * 24)
}

func aggrToPair(m []*model.AggrScore) []redis.Pair {
	r := make([]redis.Pair, len(m)+1, len(m)+1)
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
	r := make([]*pb.ItemScore, len(m), len(m))
	for i := 0; i < len(m); i++ {
		r[i] = &pb.ItemScore{
			Id:    m[i].Sid,
			Score: m[i].Score,
		}
	}
	return r
}
