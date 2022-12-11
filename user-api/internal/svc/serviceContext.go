package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/model"
	"zero-demo/user-api/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
