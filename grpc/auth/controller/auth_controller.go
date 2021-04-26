package controller

import (
	"auth/auth/service"
	"auth/grpc/auth/pb"
	"context"
)

type AuthController struct {}

func (auth *AuthController) IsAuth(_ context.Context, req *pb.Req) (rep *pb.Rep, err error) {
	e := service.GetEnforcer()
	success,err := e.Enforce(req.Sub, req.Obj, req.Act)
	return &pb.Rep{Success: success},err
}