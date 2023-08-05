package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
)

type InstanceLinkParseRequest struct {
	Token string `json:"token"`
}

func (ConsoleApi) InstanceLinkParseView(ctx *gin.Context) {
	var cr InstanceLinkParseRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

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

	res.OkWithData(gin.H{"template": templateModel.Content}, ctx)
}
