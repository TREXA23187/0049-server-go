package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"bytes"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (FileApi) FileInfoView(ctx *gin.Context) {

	filePath := ctx.Query("file_path")
	if filePath == "" {
		res.FailWithCode(res.ArgumentError, ctx)
	}

	// Send GET request to OSS URL
	resp, err := http.Get(filePath)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	reader := csv.NewReader(bytes.NewReader(data))

	header, err := reader.Read()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
	}

	res.OkWithData(header, ctx)
}
