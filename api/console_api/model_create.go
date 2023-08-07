package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/utils/jwts"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ModelCreateRequest struct {
	Name            string   `json:"name" binding:"required" msg:"Please enter name"`
	Type            string   `json:"type" binding:"required" msg:"Please enter type"`
	CodeFile        string   `json:"code_file"`
	ModelFileNames  []string `json:"model_file_names"`
	ModelFilePaths  []string `json:"model_file_paths"`
	IsGithub        bool     `json:"is_github"`
	GithubLink      string   `json:"github_link"`
	HyperParameters string   `json:"hyper_parameters"`
}

func (ConsoleApi) ModelCreateView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

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
	modelModel.Type = cr.Type
	modelModel.CreateUser = claims.UserID
	modelModel.ModelFileName = cr.ModelFileNames[0]
	modelModel.ModelFilePath = cr.ModelFilePaths[0]
	modelModel.HyperParameters = cr.HyperParameters

	if cr.IsGithub {
		modelModel.IsGithub = cr.IsGithub
		modelModel.GithubLink = cr.GithubLink
	}

	for _, v := range cr.HyperParameters {
		fmt.Println(v)
	}

	err = global.DB.Create(&modelModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.OkWithMessage("create model successfully", ctx)
}
