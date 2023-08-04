package users_api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (UserApi) MockView(ctx *gin.Context) {
	var data map[string]interface{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
