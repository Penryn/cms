package user

import (
	"cms/internal/service"
	"cms/pkg/utils"

	"github.com/gin-gonic/gin"
)


type RegisterData struct {
	Username        string `json:"username" binding:"required"`
	Sex             string `json:"sex" binding:"required"`
	PhoneNum        string `json:"phone_num" binding:"required"`
	Major           string `json:"major" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}


func Reg(c *gin.Context) {
	var data RegisterData
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.JsonErrorResponse(c, 200501, "参数错误")
		return
	}
	if data.Password != data.ConfirmPassword {
		utils.JsonErrorResponse(c, 200502, "两次密码不一致")
		return
	}
	_, err := service.GetUserByPhoneNum(data.PhoneNum)
	if err == nil {
		utils.JsonErrorResponse(c, 200503, "用户名已存在")
		return
	}
	err = service.CreateUser(data.Username, data.Sex, data.PhoneNum, data.Major, data.Password)
	if err != nil {
		utils.JsonErrorResponse(c, 200504, "注册失败")
		return
	}
	utils.JsonSuccess(c, nil)

}