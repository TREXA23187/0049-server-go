package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func (FileApi) FileUploadView(ctx *gin.Context) {

	var fileDist string
	fileType := ctx.Query("type")
	if ctype.FileType(fileType) == ctype.DataFileType {
		fileDist = "data_file"
	} else if ctype.FileType(fileType) == ctype.ModelFileType {
		fileDist = "model"
	}

	file, _ := ctx.FormFile("file") // get file from request

	dist := filepath.Join("uploads", fileDist, file.Filename)

	// save file
	err := ctx.SaveUploadedFile(file, dist)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("upload file failed", ctx)
	} else {
		res.OkWithData(gin.H{"file_path": dist}, ctx)
	}

}
