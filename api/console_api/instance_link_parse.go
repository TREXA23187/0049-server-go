package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

type InstanceLinkParseRequest struct {
	Instance string `json:"instance"`
	Token    string `json:"token"`
}

func (ConsoleApi) InstanceLinkParseView(ctx *gin.Context) {
	var cr InstanceLinkParseRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	if cr.Token != "" {
		linkClaims, err := jwts.ParseLinkToken(cr.Token)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
			return
		}

		var instanceModel models.InstanceModel
		err = global.DB.Take(&instanceModel, "id = ?", linkClaims.InstanceId).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("instance dose not exist", ctx)
			return
		}

		var templateModel models.TemplateModel
		err = global.DB.Take(&templateModel, "id = ?", linkClaims.TemplateId).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("template dose not exist", ctx)
			return
		}

		res.OkWithData(templateModel.Content, ctx)
	} else {
		var instanceModel models.InstanceModel
		err := global.DB.Take(&instanceModel, "name = ?", cr.Instance).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage("instance dose not exist", ctx)
			return
		}

		if !instanceModel.IsPublic {
			global.Log.Error(instanceModel.Name + " permission error")
			res.FailWithMessage(instanceModel.Name+" permission error", ctx)
		} else {
			var imageModel models.ImageModel
			global.DB.Take(&imageModel, "repository = ?", instanceModel.Image)

			var taskModel models.TaskModel
			global.DB.Take(&taskModel, "name = ?", imageModel.Task)

			var templateModel models.TemplateModel
			err := global.DB.Take(&templateModel, "name = ?", taskModel.Template).Error
			if err != nil {
				global.Log.Error(err)
				res.FailWithMessage("template dose not exist", ctx)
				return
			}

			res.OkWithData(templateModel.Content, ctx)
		}
	}
}
