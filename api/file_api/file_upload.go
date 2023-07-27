package file_api

import (
	"0049-server-go/global"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/plugins/qiniu"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func (FileApi) FileUploadView(ctx *gin.Context) {

	var filePrefix string
	fileType := ctx.Query("type")

	if ctype.FileType(fileType) == ctype.DataFileType {
		filePrefix = "data_file"
	} else if ctype.FileType(fileType) == ctype.ModelFileType {
		filePrefix = "model"
	}

	file, header, err := ctx.Request.FormFile("file") // get file from request
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// upload file
	filePath, err := qiniu.UploadFile(data, header.Filename, filePrefix)

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("upload file failed", ctx)
	} else {
		res.OkWithData(gin.H{"file_path": filePath}, ctx)
	}

	//err := ctx.SaveUploadedFile(file, dist)
	//if err != nil {
	//	global.Log.Error(err)
	//	res.FailWithMessage("upload file failed", ctx)
	//} else {
	//	res.OkWithData(gin.H{"file_path": dist}, ctx)
	//}

}
