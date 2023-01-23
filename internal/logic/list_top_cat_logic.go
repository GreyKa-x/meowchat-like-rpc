package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTopCatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTopCatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTopCatLogic {
	return &ListTopCatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTopCatLogic) ListTopCat(in *pb.ListTopCatReq) (*pb.ListTopCatResp, error) {
	cats, err := l.svcCtx.CatPopularityModel.ListTopK(l.ctx, in.K)
	if err != nil {
		return nil, err
	}
	return &pb.ListTopCatResp{Cats: toCatPop(cats)}, nil
}
