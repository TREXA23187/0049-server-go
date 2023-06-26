package users_api

import (
	"0049-server-go/models/res"
	pd "0049-server-go/proto"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"reflect"
)

const (
	address = "localhost:8082"
)

func (UserApi) PbListView(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pd.NewFormatDataClient(conn)

	r, err := c.DoFormat(context.Background(), &pd.Actionrequest{Text: "test", Corpus: pd.Actionrequest_NEWS})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	fmt.Println(r.Age)
	fmt.Println(reflect.TypeOf(r.Result))
	for k, v := range r.Result {
		fmt.Println(k, v)
		fmt.Println(v.Snippets)
	}

	res.OkWithMessage("1", ctx)
}
