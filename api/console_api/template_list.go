package console_api

import (
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"github.com/gin-gonic/gin"
)

func (ConsoleApi) TemplateListView(ctx *gin.Context) {
	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, _ := common.ComList(models.TemplateModel{}, common.Option{
		PageInfo: page,
	})

	res.OkWithList(list, count, ctx)
}
