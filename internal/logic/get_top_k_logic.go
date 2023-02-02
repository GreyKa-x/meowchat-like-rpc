package logic

import (
	"context"
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
	data, err := l.svcCtx.ScoreModel.ListTopK(l.ctx, in.K, int(in.Range), in.Type)
	if err != nil {
		return nil, err
	}
	return &pb.GetTopKResp{Item: data}, nil
}
