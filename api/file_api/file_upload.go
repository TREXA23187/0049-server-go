package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func (FileApi) FileUploadView(ctx *gin.Context) {

	file, _ := ctx.FormFile("file") // 从请求中获取文件
	fmt.Println(33, file.Filename)

	dist := filepath.Join("uploads/data_file", file.Filename)
	fmt.Println(dist)

	// 保存文件到指定路径
	err := ctx.SaveUploadedFile(file, dist)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("upload file failed", ctx)
	} else {
		res.OkWithNone(ctx)
	}

}
