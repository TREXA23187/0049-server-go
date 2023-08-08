package console_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/ctype"
	"0049-server-go/models/res"
	"0049-server-go/services/common"
	"0049-server-go/utils/jwts"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (ConsoleApi) ImageListView(ctx *gin.Context) {
	_claims, _ := ctx.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var page models.PageInfo
	if err := ctx.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, ctx)
		return
	}

	list, count, _ := common.ComList(models.ImageModel{}, common.Option{
		PageInfo:   page,
		CreateUser: claims.UserID,
	})

	for i, v := range list {
		if list[i].Status != ctype.ImageStatusBeingBuilt {

			var instance models.InstanceModel
			if err := global.DB.First(&instance, "image = ?", v.Repository).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					list[i].Status = ctype.ImageStatusUnused
				}
			} else {
				list[i].Status = ctype.ImageStatusInUse
			}

		}
	}

	res.OkWithList(list, count, ctx)
}
