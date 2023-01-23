package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddScoreLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddScoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddScoreLogic {
	return &AddScoreLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddScoreLogic) AddScore(in *pb.AddScoreReq) (*pb.AddScoreResp, error) {

	key := "score_" + in.Type

	val, err := l.svcCtx.Redis.ZincrbyCtx(l.ctx, key, in.Val, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.AddScoreResp{Val: val}, nil
}
