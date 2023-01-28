package logic

import (
	"context"
	"fmt"
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

// GetTopK
/*
函数逻辑：
Range == 1 时，直接从 Redis 访问当日数据
Range > 1 时，拼出 cacheKey 在 Redis 里取数据，若不存在，访问数据库提取数据入 Redis。
访问 Redis 前进行分段加锁，把 cacheKey hash 到一个数，作为 id ，对此 id 的锁加锁，
目的是过滤重复请求，相同请求同一时刻只有一个打入数据库，同时分段数量限制了最大数据库并发量，顺带缓解缓存雪崩问题
*/
func (l *GetTopKLogic) GetTopK(in *pb.GetTopKReq) (*pb.GetTopKResp, error) {

	if in.Range > 1 {
		// 长期数据
		cacheKey := fmt.Sprintf("cache:score_%s_%d_%d", in.Type, in.Range, in.K)
		Lock.lock(cacheKey)
		defer Lock.unlock(cacheKey)
		d, err := l.svcCtx.Redis.ZrangeWithScoresCtx(l.ctx, cacheKey, 0, -1)
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
			_, err = l.svcCtx.Redis.ZaddsCtx(l.ctx, cacheKey, aggrToPair(data)...)
			if err != nil {
				return nil, err
			}
			err = l.svcCtx.Redis.Expire(cacheKey, getExpireTime(in.Range))
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
