package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	pb "0049-server-go/proto"
	"0049-server-go/services/file_service"
	"0049-server-go/services/redis_service"
	"0049-server-go/utils/jwts"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

type ImageCreateRequest struct {
	Repository string `json:"repository" binding:"required" msg:"Please enter repository"`
	Tag        string `json:"tag"`
	Task       string `json:"task"`
}

func (ConsoleApi) ImageCreateView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

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

	var modelModel models.ModelModel
	if taskModel.Type == ctype.TrainingTask {
		err = global.DB.Take(&modelModel, "name = ?", taskModel.Model).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("model does not exist", ctx)
			return
		}
	}

	err = redis_service.StartBuildingImage(cr.Repository, time.Minute*10)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	imageModel.Task = cr.Task
	imageModel.Repository = cr.Repository
	imageModel.Tag = cr.Tag
	imageModel.Status = ctype.ImageStatusBeingBuilt
	imageModel.CreateUser = claims.UserID

	err = global.DB.Create(&imageModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("The image is being built", ctx)

	go func() {
		if taskModel.Type == ctype.TrainingTask {
			dataFileBytes, err := file_service.GetQiNiuFileBytes(taskModel.DataFilePath)
			modelFileBytes, err := file_service.GetQiNiuFileBytes(modelModel.ModelFilePath)

			conn := pb.NewInstanceServiceClient(global.GRPC)

			imageConfig := &pb.ImageConfig{TaskId: strconv.Itoa(int(taskModel.ID)), Type: string(taskModel.Type), Model: modelModel.Name, DataLabel: taskModel.TrainingDataLabel, TargetLabel: taskModel.TrainingLabel, Template: "", Title: fmt.Sprintf("%s: %s", string(taskModel.Type), modelModel.Name)}
			r, err := conn.CreateTrainingImage(context.Background(), &pb.CreateTrainingImageRequest{Repository: cr.Repository, Tag: cr.Tag, DataFile: dataFileBytes, ModelFile: modelFileBytes, ImageConfig: imageConfig})
			if err != nil {
				global.Log.Error("could not greet: %v", err)
				res.FailWithMessage("could not greet: %v", ctx)
			}

			var builtImageModel models.ImageModel
			global.DB.Take(&builtImageModel, "repository = ?", cr.Repository)

			builtImageModel.Status = ctype.ImageStatusUnused
			builtImageModel.ImageID = r.ImageId
			builtImageModel.Size = r.ImageSize

			global.DB.Save(&builtImageModel)

			err = redis_service.FinishBuildingImage(cr.Repository)
			if err != nil {
				log.Println("fail to finish building", err)
				return
			}

			res.OkWithMessage("create image successfully", ctx)
		}

		if taskModel.Type == ctype.DeploymentTask {
			trainedModelFileBytes, err := file_service.GetQiNiuFileBytes(taskModel.TrainedModelFilePath)

			conn := pb.NewInstanceServiceClient(global.GRPC)

			imageConfig := &pb.ImageConfig{TaskId: strconv.Itoa(int(taskModel.ID)), Type: string(taskModel.Type), Model: "", DataLabel: "", TargetLabel: "", Template: "", Title: ""}
			r, err := conn.CreateDeploymentImage(context.Background(), &pb.CreateDeploymentImageRequest{Repository: cr.Repository, Tag: cr.Tag, Template: "", TrainedModelFile: trainedModelFileBytes, ImageConfig: imageConfig})
			if err != nil {
				global.Log.Error("could not greet: %v", err)
				res.FailWithMessage("could not greet: %v", ctx)
			}

			var builtImageModel models.ImageModel
			global.DB.Take(&builtImageModel, "repository = ?", cr.Repository)

			builtImageModel.Status = ctype.ImageStatusUnused
			builtImageModel.ImageID = r.ImageId
			builtImageModel.Size = r.ImageSize

			global.DB.Save(&builtImageModel)

			err = redis_service.FinishBuildingImage(cr.Repository)
			if err != nil {
				log.Println("fail to finish building", err)
				return
			}

			res.OkWithMessage("create image successfully", ctx)
		}
	}()

}
