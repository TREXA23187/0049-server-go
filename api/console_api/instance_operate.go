package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OperationType string

const (
	StartOperation  = "start"
	StopOperation   = "stop"
	RemoveOperation = "remove"
)

type InstanceOperateRequest struct {
	InstanceID string `json:"instance_id" binding:"required" msg:"Please enter instance id"`
	Operation  string `json:"operation" binding:"required" msg:"Please enter operation"`
}

func (ConsoleApi) InstanceOperateView(ctx *gin.Context) {
	var cr InstanceOperateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
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

	r, err := c.OperateInstance(context.Background(), &pb.OperateInstanceRequest{InstanceId: cr.InstanceID, Operation: string(cr.Operation)})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	if !r.Success {
		res.FailWithMessage(fmt.Sprintf("Failed to %s the instance", string(cr.Operation)), ctx)
	}

	res.OkWithMessage(fmt.Sprintf("%s the instance successfully", string(cr.Operation)), ctx)
}
