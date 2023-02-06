package svc

import (
	"github.com/xh-polaris/meowchat-like-rpc/internal/common"
	"github.com/xh-polaris/meowchat-like-rpc/internal/config"
	"github.com/xh-polaris/meowchat-like-rpc/internal/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config
	model.LikeModel
	model.ScoreModel
	model.CatPopularityModel
	*redis.Redis
	MsgQ common.MsgQ
}

func NewServiceContext(c config.Config) *ServiceContext {
	cache := c.Redis.NewRedis()
	return &ServiceContext{
		Config:             c,
		LikeModel:          model.NewLikeModel(c.Mongo.URL, c.Mongo.DB, model.LikeCollectionName, c.CacheConf),
		ScoreModel:         model.NewScoreModel(c.Mongo.URL, c.Mongo.DB, model.ScoreCollectionName, cache),
		CatPopularityModel: model.NewCatPopularityModel(c.Mongo.URL, c.Mongo.DB, model.PopularityCollectionName),
		Redis:              cache,
		MsgQ:               common.NewMsgQImpl(c.MqConf.NameServer, c.MqConf.Retry, c.MqConf.GroupName),
	}
}
