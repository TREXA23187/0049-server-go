package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
)

type ModelCreateRequest struct {
	Name       string `json:"name" binding:"required" msg:"Please enter name"`
	CodeFile   string `json:"code_file"`
	IsGithub   bool   `json:"is_github"`
	GithubLink string `json:"github_link"`
}

func (ConsoleApi) ModelCreateView(ctx *gin.Context) {
	var cr ModelCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	var modelModel models.ModelModel
	err := global.DB.Take(&modelModel, "name = ?", cr.Name).Error
	if err == nil {
		global.Log.Error(err)
		res.FailWithMessage("model name already exists", ctx)
		return
	}

	modelModel.Name = cr.Name
	if cr.IsGithub {
		modelModel.IsGithub = cr.IsGithub
		modelModel.GithubLink = cr.GithubLink
	}

	err = global.DB.Create(&modelModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create model successfully", ctx)
}
