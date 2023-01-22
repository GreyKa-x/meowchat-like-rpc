package logic

import (
	"context"
	"fmt"
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DailyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDailyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DailyUpdateLogic {
	return &DailyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DailyUpdateLogic) DailyUpdate(in *pb.DailyUpdateReq) (*pb.DailyUpdateResp, error) {

	// 获取所有类型的zset名称
	zkeys, err := l.svcCtx.Redis.KeysCtx(l.ctx, "score_*")
	if err != nil {
		return nil, err
	}
	fmt.Println(zkeys)
	// 遍历所有zset, 对每个zset，先插入mongo再删redis
	for _, zkey := range zkeys {
		t, err := l.svcCtx.Redis.ZrangeWithScoresCtx(l.ctx, zkey, 0, -1)
		if err != nil {
			return nil, err
		}
		tname := zkey[6:]
		data := make([]model.Score, len(t), len(t))
		for i := 0; i < len(data); i++ {
			data[i].Score = t[i].Score
			data[i].Sid = t[i].Key
			data[i].Type = tname
		}
		err = l.svcCtx.ScoreModel.InsertMany(l.ctx, data)
		if err != nil {
			return nil, err
		}

		_, err = l.svcCtx.Redis.DelCtx(l.ctx, zkey)
		if err != nil {
			return nil, err
		}
	}

	return &pb.DailyUpdateResp{}, nil
}

//
