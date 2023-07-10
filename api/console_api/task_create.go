package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

type TaskCreateRequest struct {
	Name          string `json:"name" binding:"required" msg:"Please enter name"`
	Type          string `json:"type"`
	InstanceId    uint   `json:"instance_id"`
	TrainingLabel string `json:"training_label"`
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
	taskModel.InstanceID = cr.InstanceId
	taskModel.Status = ctype.StatusExited

	if ctype.TaskType(cr.Type) == ctype.TrainingTask {
		taskModel.TrainingLabel = cr.TrainingLabel
	}

	err = global.DB.Create(&taskModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create task successfully", ctx)
}
