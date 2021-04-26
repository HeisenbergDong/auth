package start

import (
	"auth/global"
	"auth/grpc/auth/controller"
	"auth/grpc/auth/pb"
	"auth/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	grpcAddress string
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{
		grpcAddress: global.CONFIG.Grpc.AuthAddress,
	}
}

func (g *GrpcServer) Run() {
	go func() {
		defer utils.RecoverPanic()
		listen, err := net.Listen("tcp", g.grpcAddress)
		if err != nil {
			global.LOG.Error("grpc server", zap.Any(" Failed to listen:", err))
		}
		s := grpc.NewServer()
		// 服务注册
		pb.RegisterAUTHServer(s, &controller.AuthController{})

		global.LOG.Info("Listen on " + g.grpcAddress)

		if err := s.Serve(listen); err != nil {
			global.LOG.Error("grpc server", zap.Any("Failed to serve:", err))
		}
	}()
}
