package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"os"
)

func (FileApi) FileInfoView(ctx *gin.Context) {

	filePath := ctx.Query("file_path")
	if filePath == "" {
		res.FailWithCode(res.ArgumentError, ctx)
	}

	file, err := os.Open(filePath)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
	}

	res.OkWithData(header, ctx)
}
