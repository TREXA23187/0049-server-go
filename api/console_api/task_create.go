package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

type TaskCreateRequest struct {
	Name          string   `json:"name" binding:"required" msg:"Please enter task name"`
	Type          string   `json:"type" binding:"required" msg:"Please enter task type"`
	Model         string   `json:"model"`
	DataFileNames []string `json:"data_file_names"`
	TrainingLabel string   `json:"training_label"`
}

func (ConsoleApi) TaskCreateView(ctx *gin.Context) {
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
	taskModel.Status = ctype.StatusExited

	if ctype.TaskType(cr.Type) == ctype.TrainingTask {
		taskModel.Model = cr.Model
		taskModel.TrainingLabel = cr.TrainingLabel

		var dataFileName, dataFilePath string
		var dataFilePaths []string

		if len(cr.DataFileNames) > 0 {

			dataFilePaths = make([]string, len(cr.DataFileNames))

			for i, fileName := range cr.DataFileNames {
				dist := filepath.Join("uploads/data_file", fileName)
				_, err := os.Stat(dist)
				if err != nil {
					if os.IsNotExist(err) {
						global.Log.Error(errors.New(fmt.Sprintf("File does not exist: %s", dist)))
						res.FailWithMessage(fmt.Sprintf("File does not exist: %s", dist), ctx)
					} else {
						global.Log.Error(errors.New(fmt.Sprintf("An error occurred while checking the file: %s", err)))
						res.FailWithMessage(fmt.Sprintf("An error occurred while checking the file: %s", err), ctx)
					}
				}

				dataFilePaths[i] = dist
			}

			dataFilePath = dataFilePaths[0]
			dataFileName = cr.DataFileNames[0]
		} else {
			dataFilePath = ""
			dataFileName = ""
		}

		taskModel.DataFileName = dataFileName
		taskModel.DataFilePath = dataFilePath

	}

	err = global.DB.Create(&taskModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create task successfully", ctx)
}
