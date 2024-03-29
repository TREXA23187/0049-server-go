package console_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	pb "0049-server-go/proto"
	"context"
	"errors"
)

func (ConsoleService) GetMaxPort() int {
	var maxPort int
	err := global.DB.Table("instance_models").Select("max(port)").Row().Scan(&maxPort)
	if err != nil {
		return 15005
	}

	return maxPort
}

func (ConsoleService) CreateInstance(instanceId, instanceName, name, description, image, url, ip string, status ctype.Status, port int, createUser uint) (*models.InstanceModel, error) {

	// Check if the user exists
	var instanceModel models.InstanceModel
	err := global.DB.Take(&instanceModel, "name = ?", name).Error
	if err == nil {
		return nil, errors.New("title already exists")
	}

	instanceModel = models.InstanceModel{
		InstanceID:   instanceId,
		InstanceName: instanceName,
		Name:         name,
		Description:  description,
		Image:        image,
		URL:          url,
		IP:           ip,
		Port:         port,
		Status:       status,
		CreateUser:   createUser,
	}

	// save to database
	err = global.DB.Create(&instanceModel).Error
	if err != nil {
		global.Log.Error(err)
		return nil, err
	}

	return &instanceModel, nil
}

func (ConsoleService) GetInstanceStatus(InstanceID string) (*pb.InstanceInfoResponse, error) {
	c := pb.NewInstanceServiceClient(global.GRPC)

	r, err := c.GetInstanceInfo(context.Background(), &pb.InstanceInfoRequest{InstanceId: InstanceID})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		return nil, err
	}

	return r, nil
}
