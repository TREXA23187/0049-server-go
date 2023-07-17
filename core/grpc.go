package core

import (
	"0049-server-go/global"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcConn *grpc.ClientConn

func ConnectGRPC() *grpc.ClientConn {
	grpcConf := global.Config.GRPC

	grpcConn, err := grpc.Dial(grpcConf.Addr(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
	}

	return grpcConn
}
