package users_api

import (
	"0049-server-go/global"
	"0049-server-go/models"
	"0049-server-go/models/res"
	"0049-server-go/plugins/log_stash"
	"0049-server-go/services/user_service"
	"0049-server-go/utils/jwts"
	"0049-server-go/utils/pwd"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"Please enter username"`
	Password string `json:"password" binding:"required" msg:"Please enter password"`
}

func (UserApi) UserLoginView(ctx *gin.Context) {
	var cr UserLoginRequest
	err := ctx.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	log := log_stash.NewLogByGin(ctx)

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		global.Log.Warn("username does not exist")
		log.Warn(fmt.Sprintf("%s username does not exist", cr.UserName))
		res.FailWithMessage("Incorrect username or password", ctx)
		return
	}

	// Verify password
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("Incorrect username or password")
		log.Warn(fmt.Sprintf("%s %s Incorrect username or password", cr.UserName, cr.Password))
		res.FailWithMessage("Incorrect username or password", ctx)
		return
	}

	roleModel, err := user_service.GetRoleByUserId(strconv.Itoa(int(userModel.ID)))
	if err != nil {
		res.FailWithValidMsg(err, &cr, ctx)
		return
	}

	fmt.Println(roleModel)

	// Successfully logged in and generate token
	token, err := jwts.GenerateToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     roleModel.RoleType,
		UserID:   userModel.ID,
	})

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("Token generation failed", ctx)
		return
	}

	res.OkWithData(token, ctx)
}
