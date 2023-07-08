package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"0049-server-go/services/console_service"
	"github.com/gin-gonic/gin"
)

type InstanceListResponse struct {
	models.InstanceModel
	Template string `json:"template"`
	Model    string `json:"model"`
}

func (ConsoleApi) InstanceListView(ctx *gin.Context) {
	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	var instances []models.InstanceModel
	global.DB.Find(&instances)

	list, count, _ := common.ComList(models.InstanceModel{}, common.Option{
		PageInfo: page,
	})

	resultList := make([]InstanceListResponse, len(list))
	for i, l := range list {
		r, err := console_service.ConsoleService{}.GetInstanceStatus(l.InstanceID)
		if err != nil {
			res.FailWithMessage(err.Error(), ctx)
		}
		list[i].Status = r.Status

		template, err := console_service.GetTemplateById(l.TemplateID)
		if err != nil {
			res.FailWithMessage(err.Error(), ctx)
		}

		model, err := console_service.GetModelById(l.ModelID)
		if err != nil {
			res.FailWithMessage(err.Error(), ctx)
		}

		resultList[i] = InstanceListResponse{
			list[i],
			template.Title,
			model.Name,
		}

	}

	res.OkWithList(resultList, count, ctx)
}
