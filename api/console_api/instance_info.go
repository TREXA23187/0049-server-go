package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InstanceInfoRequest struct {
	InstanceID string `json:"instance_id" binding:"required" msg:"Please enter instance id"`
}

func (ConsoleApi) InstanceInfoView(ctx *gin.Context) {
	var cr InstanceInfoRequest
	if err := ctx.ShouldBindQuery(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewInstanceServiceClient(conn)

	r, err := c.GetInstanceInfo(context.Background(), &pb.InstanceInfoRequest{InstanceId: cr.InstanceID})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	res.OkWithData(r, ctx)
}
