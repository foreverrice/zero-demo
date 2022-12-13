package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"

	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// metadata跨服务传输
	if md, ok := metadata.FromIncomingContext(l.ctx); ok {
		tmp := md.Get("username")
		fmt.Printf("tmp: %+v \n", tmp)
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	fmt.Println("GetUserInfo come in")
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResp{
		Id:       user.Id,
		Nickname: user.Nickname,
	}, nil
}
