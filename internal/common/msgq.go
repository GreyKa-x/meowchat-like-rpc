package common

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/xh-polaris/meowchat-like-rpc/errorx"
	"github.com/xh-polaris/meowchat-like-rpc/like"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ MsgQ = (*MsgQImpl)(nil)

type (
	MsgQ interface {
		SendCreateAsync(msg like.LikeMsg) error
		SendDeleteAsync(msg like.LikeMsg) error
	}
	MsgQImpl struct {
		producer rocketmq.Producer
	}
)

func callback(ctx context.Context, result *primitive.SendResult, err error) {
	if err != nil {
		logx.Error(err)
	}
}

func (m *MsgQImpl) SendDeleteAsync(msg like.LikeMsg) error {
	data, err := msg.Encode()
	if err != nil {
		return errorx.ErrMsgEncoder
	}
	err = m.producer.SendAsync(context.Background(), callback, &primitive.Message{Body: data, Topic: like.DeleteLikeTopic})
	if err != nil {
		return errorx.ErrMsgQ
	}
	return nil
}

func (m *MsgQImpl) SendCreateAsync(msg like.LikeMsg) error {
	data, err := msg.Encode()
	if err != nil {
		return err
	}
	err = m.producer.SendAsync(context.Background(), callback, &primitive.Message{Body: data, Topic: like.CreateLikeTopic})
	if err != nil {
		return err
	}
	return nil
}

func NewMsgQImpl(nameServer []string, retry int, groupName string) *MsgQImpl {
	mq, err := rocketmq.NewProducer(
		producer.WithNameServer(nameServer),
		producer.WithRetry(retry),
		producer.WithGroupName(groupName),
	)
	if err != nil {
		logx.Error(errorx.ErrMsgQ)
	}
	err = mq.Start()
	if err != nil {
		logx.Error(errorx.ErrMsgQ)
	}
	return &MsgQImpl{producer: mq}
}
