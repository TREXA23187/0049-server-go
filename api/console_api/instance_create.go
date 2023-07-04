package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/services/console_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type InstanceCreateRequest struct {
	Title       string `json:"title" binding:"required" msg:"Please enter title"`
	Description string `json:"description" binding:"required" msg:"Please enter description"`
	Template    string `json:"template"`
	Model       string `json:"model"`
	URL         string `json:"url"`
	DataFile    string `json:"data_file"`
}

func (ConsoleApi) InstanceCreateView(ctx *gin.Context) {
	var cr InstanceCreateRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	fmt.Println(cr)

	var templateModel models.TemplateModel
	err := global.DB.Take(&templateModel, "name = ?", cr.Template).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	var modelModel models.ModelModel
	err = global.DB.Take(&modelModel, "model_key = ?", cr.Model).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	instanceModel, err := console_service.ConsoleService{}.CreateInstance(cr.Title, cr.Description, strconv.Itoa(int(templateModel.ID)), strconv.Itoa(int(modelModel.ID)), cr.URL, cr.DataFile)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	res.Ok(gin.H{"instance_id": instanceModel.ID}, "create instance successfully", ctx)
}
