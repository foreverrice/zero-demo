package user

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"zero-demo/model"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCreateLogic {
	return &UserCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCreateLogic) UserCreate(req *types.UserCreateReq) (resp *types.UserCreateResp, err error) {
	if err := l.svcCtx.UserModel.TransCtx(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		user := &model.User{}
		user.Mobile = req.Mobile
		user.Nickname = req.Nickname

		// 添加user
		dbResult, err := l.svcCtx.UserModel.TransInsert(ctx, session, user)
		if err != nil {
			return err
		}
		userId, _ := dbResult.LastInsertId()

		// 添加userData
		userData := &model.UserData{}
		userData.UserId = userId
		userData.Data = "xxxx"
		if _, err := l.svcCtx.UserDataModel.TransInsert(ctx, session, userData); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &types.UserCreateResp{
		Flag: true,
	}, nil
}
