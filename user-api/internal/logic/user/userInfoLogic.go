package user

import (
	"context"
	"errors"
	"zero-demo/model"

	"zero-demo/user-api/internal/svc"
	"zero-demo/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)

	if err != nil && err != model.ErrNotFound {
		return nil, errors.New("查询数据失败")
	}

	if user == nil {
		return nil, errors.New("用户不存在")
	}

	return &types.UserInfoResp{
		UserId:   user.Id,
		Nickname: user.Nickname,
	}, nil
}
