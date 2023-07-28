package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func (FileApi) FileDownloadView(ctx *gin.Context) {

	filePath := ctx.Query("file_path")

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

	ctx.Data(http.StatusOK, resp.Header.Get("Content-Type"), data)
}
