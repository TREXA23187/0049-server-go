package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"time"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (UserApi) UploadFileGrpcView(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewFileServiceClient(conn)

	file, err := os.Open("data/test.txt")
	if err != nil {
		global.Log.Error("failed to open file: ", err)
		res.FailWithMessage("failed to open file: ", ctx)
	}
	defer file.Close()

	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	
	stream, err := c.UploadFile(ctx)
	if err != nil {
		global.Log.Error("failed to upload file", err)
		res.FailWithMessage("failed to upload file", ctx)
	}

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
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
