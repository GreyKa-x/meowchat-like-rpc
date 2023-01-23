// Code generated by goctl. DO NOT EDIT!
// Source: like.proto

package server

import (
	"context"

	"github.com/xh-polaris/meowchat-like-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

// 点赞/取消赞
func (s *LikeServer) DoLike(ctx context.Context, in *pb.DoLikeReq) (*pb.DoLikeResp, error) {
	l := logic.NewDoLikeLogic(ctx, s.svcCtx)
	return l.DoLike(in)
}

// 获取用户是否点赞
func (s *LikeServer) GetUserLike(ctx context.Context, in *pb.GetUserLikedReq) (*pb.GetUserLikedResp, error) {
	l := logic.NewGetUserLikeLogic(ctx, s.svcCtx)
	return l.GetUserLike(in)
}

// 获取目标点赞数
func (s *LikeServer) GetTargetLikes(ctx context.Context, in *pb.GetTargetLikesReq) (*pb.GetTargetLikesResp, error) {
	l := logic.NewGetTargetLikesLogic(ctx, s.svcCtx)
	return l.GetTargetLikes(in)
}

func (s *LikeServer) GetTopK(ctx context.Context, in *pb.GetTopKReq) (*pb.GetTopKResp, error) {
	l := logic.NewGetTopKLogic(ctx, s.svcCtx)
	return l.GetTopK(in)
}

func (s *LikeServer) AddScore(ctx context.Context, in *pb.AddScoreReq) (*pb.AddScoreResp, error) {
	l := logic.NewAddScoreLogic(ctx, s.svcCtx)
	return l.AddScore(in)
}

func (s *LikeServer) DailyUpdate(ctx context.Context, in *pb.DailyUpdateReq) (*pb.DailyUpdateResp, error) {
	l := logic.NewDailyUpdateLogic(ctx, s.svcCtx)
	return l.DailyUpdate(in)
}

// popularity
func (s *LikeServer) ListCatPopularity(ctx context.Context, in *pb.ListCatPopularityReq) (*pb.ListCatPopularityResp, error) {
	l := logic.NewListCatPopularityLogic(ctx, s.svcCtx)
	return l.ListCatPopularity(in)
}

func (s *LikeServer) AddCatPopularity(ctx context.Context, in *pb.AddCatPopularityReq) (*pb.AddCatPopularityResp, error) {
	l := logic.NewAddCatPopularityLogic(ctx, s.svcCtx)
	return l.AddCatPopularity(in)
}

func (s *LikeServer) ListTopCat(ctx context.Context, in *pb.ListTopCatReq) (*pb.ListTopCatResp, error) {
	l := logic.NewListTopCatLogic(ctx, s.svcCtx)
	return l.ListTopCat(in)
}
