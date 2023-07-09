package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

type ModelCreateRequest struct {
	Name           string   `json:"name" binding:"required" msg:"Please enter name"`
	CodeFile       string   `json:"code_file"`
	ModelFileNames []string `json:"model_file_names"`
	IsGithub       bool     `json:"is_github"`
	GithubLink     string   `json:"github_link"`
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

	var modelFileName, modelFilePath string
	var modelFilePaths []string

	if len(cr.ModelFileNames) > 0 {

		modelFilePaths = make([]string, len(cr.ModelFileNames))

		for i, fileName := range cr.ModelFileNames {
			dist := filepath.Join("uploads/model", fileName)
			_, err := os.Stat(dist)
			if err != nil {
				if os.IsNotExist(err) {
					global.Log.Error(errors.New(fmt.Sprintf("File does not exist: %s", dist)))
					res.FailWithMessage(fmt.Sprintf("File does not exist: %s", dist), ctx)
				} else {
					global.Log.Error(errors.New(fmt.Sprintf("An error occurred while checking the file: %s", err)))
					res.FailWithMessage(fmt.Sprintf("An error occurred while checking the file: %s", err), ctx)
				}
			}

			modelFilePaths[i] = dist
		}

		modelFilePath = modelFilePaths[0]
		modelFileName = cr.ModelFileNames[0]
	} else {
		modelFilePath = ""
		modelFileName = ""
	}

	modelModel.Name = cr.Name
	modelModel.ModelFileName = modelFileName
	modelModel.ModelFilePath = modelFilePath

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
