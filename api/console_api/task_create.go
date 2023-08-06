package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type TaskCreateRequest struct {
	Name                 string         `json:"name" binding:"required" msg:"Please enter task name"`
	Type                 string         `json:"type" binding:"required" msg:"Please enter task type"`
	Model                string         `json:"model"`
	Template             string         `json:"template"`
	DataFileNames        []string       `json:"data_file_names"`
	DataFilePaths        []string       `json:"data_file_paths"`
	TrainedModelFileName string         `json:"trained_model_file_name"`
	TrainedModelFilePath string         `json:"trained_model_file_path"`
	TrainingLabel        string         `json:"training_label"`
	TrainingDataLabel    []string       `json:"training_data_label"`
	EnableAdvance        bool           `json:"enable_advance"`
	HyperParameters      map[string]any `json:"hyper_parameters"`
}

func (ConsoleApi) TaskCreateView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var cr TaskCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var taskModel models.TaskModel
	err := global.DB.Take(&taskModel, "name = ?", cr.Name).Error
	if err == nil {
		global.Log.Error(err)
		res.FailWithMessage("task name already exists", ctx)
		return
	}

	taskModel.Name = cr.Name
	taskModel.Type = ctype.TaskType(cr.Type)
	taskModel.Status = ctype.TaskStatusPending
	taskModel.CreateUser = claims.UserID

	if ctype.TaskType(cr.Type) == ctype.TrainingTask {
		taskModel.Model = cr.Model
		taskModel.TrainingLabel = cr.TrainingLabel

		trainingDataLabel, err := json.Marshal(cr.TrainingDataLabel)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
			return
		}
		taskModel.TrainingDataLabel = string(trainingDataLabel)

		taskModel.DataFileName = cr.DataFileNames[0]
		taskModel.DataFilePath = cr.DataFilePaths[0]

		taskModel.EnableAdvance = cr.EnableAdvance
		if cr.EnableAdvance {
			hyperParamBytes, err := json.Marshal(cr.HyperParameters)
			if err != nil {
				global.Log.Error(err)
				res.FailWithMessage(err.Error(), ctx)
				return
			}

			taskModel.HyperParameters = string(hyperParamBytes)
		}
	}

	if ctype.TaskType(cr.Type) == ctype.DeploymentTask {
		taskModel.Template = cr.Template
		taskModel.TrainedModelFileName = cr.TrainedModelFileName
		taskModel.TrainedModelFilePath = cr.TrainedModelFilePath
	}

	err = global.DB.Create(&taskModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create task successfully", ctx)
}
