package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) InstanceInfoView(ctx *gin.Context) {
	instanceId := ctx.Query("instance_id")

	c := pb.NewInstanceServiceClient(global.GRPC)

	r, err := c.GetInstanceInfo(context.Background(), &pb.InstanceInfoRequest{InstanceId: instanceId, NeedLogs: false})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	res.OkWithData(r, ctx)
}
