package console_service

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	pb "0049-server-go/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"path/filepath"
)

func (ConsoleService) GetMaxPort() int {
	var maxPort int
	err := global.DB.Table("instance_models").Select("max(port)").Row().Scan(&maxPort)
	if err != nil {
		return 15005
	}

	return maxPort
}

func (ConsoleService) CreateInstance(instanceId, instanceName, title, description, templateId, modelId, url, ip, status string, dataFileNames []string, port int) (*models.InstanceModel, error) {

	// Check if the user exists
	var instanceModel models.InstanceModel
	err := global.DB.Take(&instanceModel, "title = ?", title).Error
	if err == nil {
		return nil, errors.New("title already exists")
	}

	dataFilePaths := make([]string, len(dataFileNames))

	for _, fileName := range dataFileNames {
		dist := filepath.Join("uploads/data_file", fileName)
		_, err := os.Stat(dist)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, errors.New(fmt.Sprintf("File does not exist: %s", dist))
			} else {
				// some other error happened
				return nil, errors.New(fmt.Sprintf("An error occurred while checking the file: %s", err))
			}
		}

		dataFilePaths = append(dataFilePaths, dist)
	}

	var dataFilePath string
	if len(dataFilePaths) > 0 {
		dataFilePath = dataFilePaths[0]
	} else {
		dataFilePath = ""
	}

	instanceModel = models.InstanceModel{
		InstanceID:   instanceId,
		InstanceName: instanceName,
		Title:        title,
		Description:  description,
		TemplateID:   templateId,
		ModelID:      modelId,
		URL:          url,
		IP:           ip,
		Port:         port,
		Status:       status,
		DataFilePath: dataFilePath,
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
	conn, err := grpc.Dial(ctype.GRPC_ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		global.Log.Error("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	c := pb.NewInstanceServiceClient(conn)

	r, err := c.GetInstanceInfo(context.Background(), &pb.InstanceInfoRequest{InstanceId: InstanceID})
	if err != nil {
		global.Log.Error("could not greet: %v", err)
		return nil, err
	}

	return r, nil
}
