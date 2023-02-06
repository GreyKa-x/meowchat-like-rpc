package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-like-rpc/errorx"
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/xh-polaris/meowchat-like-rpc/internal/svc"
	like2 "github.com/xh-polaris/meowchat-like-rpc/like"
	"github.com/xh-polaris/meowchat-like-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type DoLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDoLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DoLikeLogic {
	return &DoLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DoLike 点赞/取消赞
func (l *DoLikeLogic) DoLike(in *pb.DoLikeReq) (*pb.DoLikeResp, error) {
	// 判断是否点过赞
	getUserLikeLogic := NewGetUserLikeLogic(l.ctx, l.svcCtx)
	data := &pb.GetUserLikedReq{
		UserId:   in.UserId,
		TargetId: in.TargetId,
		Type:     in.Type,
	}
	response, _ := getUserLikeLogic.GetUserLike(data)
	switch response.Liked {
	case false:
		// 插入数据
		likeModel := l.svcCtx.LikeModel
		like := &model.Like{
			UserId:       in.UserId,
			TargetId:     in.TargetId,
			TargetType:   in.Type,
			AssociatedId: in.AssociatedId,
		}
		err := likeModel.Insert(l.ctx, like)
		if err == nil {

			msg := like2.LikeMsg{
				Id:           like.ID.Hex(),
				UserId:       like.UserId,
				TargetId:     like.TargetId,
				TargetType:   like.TargetType,
				AssociatedId: like.AssociatedId,
				Time:         like.CreateAt.Unix(),
			}
			err := l.svcCtx.MsgQ.SendCreateAsync(msg)
			if err != nil {
				logx.Error(err)
			}

			return &pb.DoLikeResp{}, nil
		} else {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
	case true:
		likeModel := l.svcCtx.LikeModel
		ID, err := likeModel.GetId(l.ctx, in.UserId, in.TargetId, in.Type)
		if err != nil {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
		err = likeModel.Delete(l.ctx, ID)
		if err == nil {

			msg := like2.LikeMsg{
				Id:           ID,
				UserId:       in.UserId,
				TargetId:     in.TargetId,
				TargetType:   in.Type,
				AssociatedId: in.AssociatedId,
				Time:         time.Now().Unix(),
			}
			err := l.svcCtx.MsgQ.SendDeleteAsync(msg)
			if err != nil {
				logx.Error(err)
			}

			return &pb.DoLikeResp{}, nil
		} else {
			return &pb.DoLikeResp{}, errorx.ErrDataBase
		}
	default:
		return &pb.DoLikeResp{}, nil
	}
}
