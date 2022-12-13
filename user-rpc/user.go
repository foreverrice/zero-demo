package main

import (
	"flag"
	"fmt"

	"zero-demo/user-rpc/internal/config"
	"zero-demo/user-rpc/internal/server"
	"zero-demo/user-rpc/internal/svc"
	"zero-demo/user-rpc/pb"

	"context"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	svr := server.NewUsercenterServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUsercenterServer(grpcServer, svr)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// 拦截器
	s.AddUnaryInterceptors(TestServerInterceptor)

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func TestServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Printf("TestServerInterceptor ===> start \n")

	fmt.Printf("req ===> %+v \n", req)
	fmt.Printf("info ===> %+v \n", info)

	resp, err = handler(ctx, req)

	fmt.Printf("TestServerInterceptor ===> end \n")

	return resp, err
}
