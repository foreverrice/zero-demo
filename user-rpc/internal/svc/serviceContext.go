package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/model"
	"zero-demo/user-rpc/internal/config"
	"zero-demo/user-rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	UserDataModel model.UserDataModel
	UserRpc       usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserDataModel: model.NewUserDataModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
