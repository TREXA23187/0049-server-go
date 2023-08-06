package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models/res"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type InterfaceDataRedirectRequest struct {
	Data map[string]interface{} `json:"data"`
	URL  string                 `json:"url" `
}

func (ConsoleApi) InterfaceDataRedirectView(ctx *gin.Context) {
	var cr InterfaceDataRedirectRequest
	if err := ctx.ShouldBindJSON(&cr); err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	byteData, err := json.Marshal(cr.Data)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}

	resp, err := http.Post("http://"+cr.URL, "application/json", bytes.NewBuffer(byteData))
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage(err.Error(), ctx)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	res.OkWithData(string(body), ctx)
}
