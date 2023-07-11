package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ImageCreateRequest struct {
	Repository string `json:"repository" binding:"required" msg:"Please enter repository"`
	Tag        string `json:"tag"`
	Task       string `json:"task"`
}

func (ConsoleApi) ImageCreateView(ctx *gin.Context) {
	var cr ImageCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, "repository = ?", cr.Repository).Error
	if err == nil {
		global.Log.Error(err)
		res.FailWithMessage("image repository already exists", ctx)
		return
	}

	var taskModel models.TaskModel
	err = global.DB.Take(&taskModel, "name = ?", cr.Task).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("task does not exist", ctx)
		return
	}

	imageModel.Task = cr.Task
	imageModel.Repository = cr.Repository
	imageModel.Tag = cr.Tag

	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewInstanceServiceClient(conn)

	r, err := c.CreateImage(context.Background(), &pb.CreateImageRequest{Repository: cr.Repository, Tag: cr.Tag})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	imageModel.ImageID = r.ImageId
	imageModel.Size = r.ImageSize

	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create image successfully", ctx)
}
