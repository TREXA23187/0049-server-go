package users_api

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
	"io"
	"log"
	"net/http"
	"time"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (UserApi) UploadFileGrpcView(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	fmt.Println(file.Filename)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		global.Log.Error("failed to open file: ", err)
		res.FailWithMessage("failed to open file: ", ctx)
	}
	defer src.Close()

	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewFileServiceClient(conn)

	cont, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	stream, err := c.UploadFile(cont)
	if err != nil {
		global.Log.Error("failed to upload file", err)
		res.FailWithMessage("failed to upload file", ctx)
	}

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err == io.EOF {
			response, err := stream.CloseAndRecv()
			if err != nil {
				log.Fatalf("failed to receive response: %v", err)
			}
			log.Printf("Response: %s", response.GetFlag())

			res.OkWithData(response.GetFlag(), ctx)
			return
		}
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}

		err = stream.Send(&pb.Bytes{Value: buf[:n]})
		if err != nil {
			log.Fatalf("failed to read file: %v", err)
		}
	}
}
