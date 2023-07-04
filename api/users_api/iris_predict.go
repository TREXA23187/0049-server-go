package users_api

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

type IrisPredictRequest struct {
	SepalLength string `json:"sepal_length" binding:"required" msg:"Please enter sepal_length"`
	SepalWidth  string `json:"sepal_width" binding:"required" msg:"Please enter sepal_width"`
	PetalLength string `json:"petal_length" binding:"required" msg:"Please enter petal_length"`
	PetalWidth  string `json:"petal_width" binding:"required" msg:"Please enter petal_width"`
}

func (UserApi) IrisPredictView(ctx *gin.Context) {
	var cr IrisPredictRequest
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
	c := pb.NewPredictionTaskClient(conn)

	r, err := c.LaunchTask(context.Background(), &pb.TaskRequest{SepalLength: cr.SepalLength, SepalWidth: cr.SepalWidth, PetalLength: cr.PetalLength, PetalWidth: cr.PetalWidth})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	res.OkWithData(r, ctx)
}
