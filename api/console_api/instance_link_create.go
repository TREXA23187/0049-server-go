package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type InstanceLinkCreateRequest struct {
	Instance string `json:"instance"`
	Template string `json:"template"`
	LinkType string `json:"link_type"`
	Autofill bool   `json:"autofill"`
}

func (ConsoleApi) InstanceLinkCreateView(ctx *gin.Context) {
	var cr InstanceLinkCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var instanceModel models.InstanceModel
	err := global.DB.Take(&instanceModel, "name = ?", cr.Instance).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("instance dose not exist", ctx)
		return
	}

	var templateModel models.TemplateModel
	err = global.DB.Take(&templateModel, "name = ?", cr.Template).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("template dose not exist", ctx)
		return
	}

	token, err := jwts.GenerateLinkToken(jwts.LinkJwtPayLoad{
		InstanceId: instanceModel.ID,
		TemplateId: templateModel.ID,
		LinkType:   cr.LinkType,
		Autofill:   cr.Autofill,
	}, time.Hour*time.Duration(global.Config.Jwt.Expires))

	res.Ok(gin.H{"link": fmt.Sprintf("%s:3000/interface/%s?token=%s", global.Config.System.Host, templateModel.Name, token)}, "create link successfully", ctx)
}
