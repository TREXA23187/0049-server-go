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
	Instance   string `json:"instance"`
	Template   string `json:"template"`
	LinkType   string `json:"link_type"`
	Expiration int    `json:"expiration"`
	Autofill   bool   `json:"autofill"`
}

type InstanceLinkCreateResponse struct {
	Link        string `json:"link"`
	Autofill    bool   `json:"autofill"`
	TemplateId  uint   `json:"template_id"`
	Redirection string `json:"redirection"`
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

	var link string
	if cr.LinkType == "private" {
		token, err := jwts.GenerateLinkToken(jwts.LinkJwtPayLoad{
			InstanceId: instanceModel.ID,
			TemplateId: templateModel.ID,
			LinkType:   cr.LinkType,
			Autofill:   cr.Autofill,
		}, time.Hour*24*time.Duration(cr.Expiration))
		if err != nil {
			global.Log.Error(err)
			res.FailWithMessage(err.Error(), ctx)
			return
		}

		link = fmt.Sprintf("%s:3000/interface/%s?token=%s", global.Config.System.Host, instanceModel.Name, token)
		instanceModel.Link = link
		instanceModel.IsPublic = false
		global.DB.Save(&instanceModel)

		var instanceLinkCreateResponse InstanceLinkCreateResponse
		instanceLinkCreateResponse.Link = link
		instanceLinkCreateResponse.Autofill = cr.Autofill
		if cr.Autofill {
			instanceLinkCreateResponse.TemplateId = templateModel.ID
			instanceLinkCreateResponse.Redirection = instanceModel.URL + "/predict"
		}

		res.Ok(instanceLinkCreateResponse, "create link successfully", ctx)
	} else {
		link = fmt.Sprintf("%s:3000/interface/%s", global.Config.System.Host, instanceModel.Name)
		instanceModel.Link = link
		instanceModel.IsPublic = true
		global.DB.Save(&instanceModel)

		var instanceLinkCreateResponse InstanceLinkCreateResponse
		instanceLinkCreateResponse.Link = link
		instanceLinkCreateResponse.Autofill = cr.Autofill
		if cr.Autofill {
			instanceLinkCreateResponse.TemplateId = templateModel.ID
			instanceLinkCreateResponse.Redirection = instanceModel.URL + "/predict"
		}

		res.Ok(instanceLinkCreateResponse, "create link successfully", ctx)
	}
}
