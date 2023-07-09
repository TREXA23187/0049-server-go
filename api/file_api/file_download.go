package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func (FileApi) FileDownloadView(ctx *gin.Context) {

	filePath := ctx.Query("file_path")
	//fileName := ctx.Query("file_name")

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			global.Log.Error(fmt.Sprintf("File does not exist: %s", filePath))
			res.FailWithMessage(fmt.Sprintf("File does not exist: %s", filePath), ctx)
		} else {
			// some other error happened
			global.Log.Error(fmt.Sprintf("An error occurred while checking the file: %s", err))
			res.FailWithMessage(fmt.Sprintf("An error occurred while checking the file: %s", err), ctx)
		}
	}

	ctx.File(filePath)

}
