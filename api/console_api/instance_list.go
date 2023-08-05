package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"0049-server-go/services/console_service"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

type InstanceListResponse struct {
	models.InstanceModel
	Task     string `json:"task"`
	TaskType string `json:"task_type"`
	Template string `json:"template"`
}

func (ConsoleApi) InstanceListView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var instances []models.InstanceModel
	global.DB.Find(&instances)

	list, count, _ := common.ComList(models.InstanceModel{}, common.Option{
		PageInfo:   page,
		CreateUser: claims.UserID,
	})

	resultList := make([]InstanceListResponse, len(list))
	for i, l := range list {
		r, err := console_service.ConsoleService{}.GetInstanceStatus(l.InstanceID)
		if err != nil {
			res.FailWithMessage(err.Error(), ctx)
		}
		list[i].Status = ctype.Status(r.Status)

		var imageModel models.ImageModel
		err = global.DB.Take(&imageModel, "repository = ?", list[i].Image).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
			return
		}

		var taskModel models.TaskModel
		err = global.DB.Take(&taskModel, "name = ?", imageModel.Task).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
			return
		}

		resultList[i] = InstanceListResponse{
			list[i],
			taskModel.Name,
			string(taskModel.Type),
			taskModel.Template,
		}
	}

	res.OkWithList(resultList, count, ctx)
}
