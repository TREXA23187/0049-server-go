package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/plugins/qiniu"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TaskOperateRequest struct {
	TaskId    string `json:"task_id" binding:"required" msg:"Please enter task id"`
	Operation string `json:"operation" binding:"required" msg:"Please enter operation"`
	ModelFile string `json:"model_file"`
}

func (ConsoleApi) TaskOperateView(ctx *gin.Context) {
	var cr TaskOperateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var task models.TaskModel
	err := global.DB.Take(&task, "id=?", cr.TaskId).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
	}

	if task.Status != ctype.TaskStatusCompleted && task.Status != ctype.TaskStatusFailed {
		if cr.Operation == "success" {
			task.Status = ctype.TaskStatusCompleted
		} else if cr.Operation == "fail" {
			task.Status = ctype.TaskStatusFailed
		}
	}

	if cr.Operation == "success" {
		trainedModelFile, err := base64.StdEncoding.DecodeString(cr.ModelFile)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
		}

		trainedModelFileName := fmt.Sprintf("task_%d.pickle", task.ID)
		trainedModelFilePath, err := qiniu.UploadFile(trainedModelFile, trainedModelFileName, "trained_model_file")
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
		}

		task.TrainedModelFileName = trainedModelFileName
		task.TrainedModelFilePath = trainedModelFilePath
	}

	global.DB.Save(&task)

	res.OkWithMessage(fmt.Sprintf("%s the task successfully", cr.Operation), ctx)
}
