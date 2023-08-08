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

func (ConsoleApi) ImageRemoveView(ctx *gin.Context) {
	imageId := ctx.Query("image_id")

	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewInstanceServiceClient(conn)

	r, err := c.DeleteImage(context.Background(), &pb.DeleteImageRequest{ImageId: imageId})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	if r.Success {
		global.DB.Delete(&models.ImageModel{}, "image_id = ?", imageId)
		res.OkWithMessage("image deleted", ctx)
	} else {
		res.FailWithMessage(r.Msg, ctx)
	}
}
