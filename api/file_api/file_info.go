package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"0049-server-go/services/file_service"
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
)

func (FileApi) FileInfoView(ctx *gin.Context) {

	filePath := ctx.Query("file_path")
	if filePath == "" {
		res.FailWithCode(res.ArgumentError, ctx)
	}

	fileBytes, err := file_service.GetQiNiuFileBytes(filePath)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	reader := csv.NewReader(bytes.NewReader(fileBytes))

	header, err := reader.Read()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
	}

	res.OkWithData(header, ctx)
}
