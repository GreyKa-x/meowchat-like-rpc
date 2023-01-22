package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCatPopularityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCatPopularityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCatPopularityLogic {
	return &AddCatPopularityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCatPopularityLogic) AddCatPopularity(in *pb.AddCatPopularityReq) (*pb.AddCatPopularityResp, error) {
	err := l.svcCtx.CatPopularityModel.AddPopularity(l.ctx, in.CatId, in.Val)
	if err != nil {
		return nil, err
	}
	return &pb.AddCatPopularityResp{}, nil
}
