package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"0049-server-go/services/console_service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"strings"
)

type InstanceCreateRequest struct {
	Title       string `json:"title" binding:"required" msg:"Please enter title"`
	Description string `json:"description" binding:"required" msg:"Please enter description"`
	Template    string `json:"template"`
	Model       string `json:"model"`
	URL         string `json:"url"`
	DataFile    string `json:"data_file"`
}

func (ConsoleApi) InstanceCreateView(ctx *gin.Context) {
	var cr InstanceCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	file, _ := ctx.FormFile("data_file")
	fmt.Println(33, file)

	var templateModel models.TemplateModel
	err := global.DB.Take(&templateModel, "title = ?", cr.Template).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	var modelModel models.ModelModel
	err = global.DB.Take(&modelModel, "name = ?", cr.Model).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		res.FailWithMessage("did not connect: %v", ctx)
	}
	defer conn.Close()
	c := pb.NewInstanceServiceClient(conn)

	var url, ip string
	var port int

	if cr.URL == "" {
		ip = "127.0.0.1"
		port = console_service.ConsoleService{}.GetMaxPort() + 1
		url = fmt.Sprintf("%s:%d", ip, port)
	} else {
		s := strings.Split(url, ":")
		ip = s[0]
		port, _ = strconv.Atoi(s[1])
	}

	r, err := c.CreateInstance(context.Background(), &pb.CreateInstanceRequest{Template: cr.Template, Model: cr.Model, Url: url})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	instanceModel, err := console_service.ConsoleService{}.CreateInstance(r.InstanceId, r.InstanceName, cr.Title, cr.Description, strconv.Itoa(int(templateModel.ID)), strconv.Itoa(int(modelModel.ID)), url, ip, "running", cr.DataFile, port)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.Ok(gin.H{"id": instanceModel.ID, "instance_id": r.InstanceId}, "create instance successfully", ctx)
}
