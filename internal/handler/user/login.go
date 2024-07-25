package user

import (
	"cms/internal/service"
	"cms/pkg/utils"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	PhoneNum string `json:"phone_num" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var data LoginData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	user, err := service.GetUserByPhoneNum(data.PhoneNum)
	if err != nil {
		utils.JsonErrorResponse(c, 200505, "用户不存在")
		return
	}
	if user.Password != data.Password {
		utils.JsonErrorResponse(c, 200506, "密码错误")
		return
	}
	utils.JsonSuccess(c, nil)
}