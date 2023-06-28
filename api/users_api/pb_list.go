package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"reflect"
)

const (
	address = "localhost:8088"
)

func (UserApi) PbListView(ctx *gin.Context) {
	//conn, err := grpc.Dial(address, grpc.WithInsecure())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewFormatDataClient(conn)

	r, err := c.DoFormat(context.Background(), &pb.Actionrequest{Text: "test", Corpus: pb.Actionrequest_NEWS})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}
	fmt.Println(r)
	fmt.Println(reflect.TypeOf(r.Result))
	for k, v := range r.Result {
		fmt.Println(k, v)
		fmt.Println(v.Snippets)
	}

	res.OkWithData(r, ctx)
}
