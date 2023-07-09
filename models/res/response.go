package res

import (
	"0049-server-go/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

type FileResponse struct {
	FileName string `json:"file_name"`
	FileData []byte `json:"file_data"`
}

const (
	Success = 0
	Error   = -1
)

func Result(code int, data any, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func Ok(data any, msg string, ctx *gin.Context) {
	Result(Success, data, msg, ctx)
}
func OkWithData(data any, ctx *gin.Context) {
	Result(Success, data, "Success", ctx)
}
func OkWithMessage(msg string, ctx *gin.Context) {
	Result(Success, map[string]any{}, msg, ctx)
}
func OkWithNone(ctx *gin.Context) {
	Result(Success, map[string]any{}, "Success", ctx)
}
func OkWithList(list any, count int64, ctx *gin.Context) {
	OkWithData(ListResponse[any]{
		Count: count,
		List:  list,
	}, ctx)
}
func OkWithFile(fileName string, fileData []byte, ctx *gin.Context) {
	ctx.Header("Content-Type", "application/octet-stream")

	OkWithData(FileResponse{
		FileName: fileName,
		FileData: fileData,
	}, ctx)
}

func Fail(data any, msg string, ctx *gin.Context) {
	Result(Error, data, msg, ctx)
}
func FailWithMessage(msg string, ctx *gin.Context) {
	Result(Error, map[string]any{}, msg, ctx)
}
func FailWithCode(code ErrorCode, ctx *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, ctx)
		return
	}

	Result(Error, map[string]any{}, "Unknown Error", ctx)
}
func FailWithValidMsg(err error, obj any, ctx *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMessage(msg, ctx)
}
