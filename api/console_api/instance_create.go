package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"0049-server-go/services/console_service"
	"0049-server-go/utils/jwts"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strconv"
	"strings"
)

type InstanceCreateRequest struct {
	Name        string `json:"name" binding:"required" msg:"Please enter name"`
	Description string `json:"description" binding:"required" msg:"Please enter description"`
	Image       string `json:"image" binding:"required" msg:"Please enter image"`
	URL         string `json:"url"`
}

func (ConsoleApi) InstanceCreateView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr InstanceCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var imageModel models.ImageModel
	err := global.DB.Take(&imageModel, "repository = ?", cr.Image).Error
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

	r, err := c.CreateInstance(context.Background(), &pb.CreateInstanceRequest{ImageName: cr.Image, Url: url})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		res.FailWithMessage("could not greet: %v", ctx)
	}

	instanceModel, err := console_service.ConsoleService{}.CreateInstance(r.InstanceId, r.InstanceName, cr.Name, cr.Description, cr.Image, url, ip, ctype.StatusRunning, port, claims.UserID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	imageModel.Status = ctype.ImageStatusInUse
	global.DB.Save(&imageModel)

	res.Ok(gin.H{"id": instanceModel.ID, "instance_id": r.InstanceId}, "create instance successfully", ctx)
}
