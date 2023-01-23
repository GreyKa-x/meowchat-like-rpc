package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCatPopularityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCatPopularityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCatPopularityLogic {
	return &ListCatPopularityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// popularity
func (l *ListCatPopularityLogic) ListCatPopularity(in *pb.ListCatPopularityReq) (*pb.ListCatPopularityResp, error) {
	popularity, err := l.svcCtx.CatPopularityModel.ListPopularity(l.ctx, in.CatId)
	if err != nil {
		return nil, err
	}
	return &pb.ListCatPopularityResp{Cats: toCatPop(popularity)}, nil
}
