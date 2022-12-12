package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"zero-demo/model"
	"zero-demo/user-api/internal/config"
	"zero-demo/user-api/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	TestMiddleware rest.Middleware
	UserModel      model.UserModel
	UserDataModel  model.UserDataModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
		UserModel:      model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel:  model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
