package controllers

import (
	"TTMS/models"
	"TTMS/service"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//有关用户的controller代码
const RootPath="./img/"
func Login(c *gin.Context){
	p:=new(models.ParamsLogin)
	err:=c.ShouldBind(&p)
	if err!=nil{
		ResponseError(c,CodeInvalidParams)
		zap.L().Error("Login ShouldBind Error")
		return
	}
	p1, err := service.Login(p)
	if p1.IsLogin == 1 && p1.IsDelete == 1 {
		ResponseErrorWithMsg(c, CodeInvalidPassword, "登录失败")
		return
	}
	if err != nil {
		if err == errors.New("登录失败") {
			ResponseErrorWithMsg(c, CodeInvalidPassword, "登录失败")
			return
		} else {
			zap.L().Error("Login Service error", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	ResponseSuccess(c, gin.H{
		"username": p1.Username,
		"token":    p1.Token,
	})
}

// Register 注册
func Register(c *gin.Context) {
	p := new(models.ParamsRegister)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("Register ShouldBind Error")
		return
	}
	err = service.Register(p)
	if err != nil {
		zap.L().Error("service.Register error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "注册成功")
}

// AddAdmin 添加管理员
func AddAdmin(c *gin.Context) {
	p := new(models.ParamsAdminmsg)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("AddAdmin ShouldBind Error")
		return
	}
	err = service.AddAdmin(p)
	if err != nil {
		zap.L().Error("service.AddAdmin error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "添加成功")
}

// GetAllMsg 获取所有用户信息
func GetAllMsg(c *gin.Context) {

	p1, err := service.GetAllMsg()

	if err != nil {
		zap.L().Error("service.GetAllMsg error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")

		return
	}
	ResponseSuccess(c, p1)
}

// GetUserMsgById 获取所有信息
func GetUserMsgById(c *gin.Context) {
	var p string
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("GetUserMsgById ShouldBind Error")
		return
	}
	p1, err := service.GetMsgById(p)

	if err != nil {
		zap.L().Error("service.GetMsgById error", zap.Error(err))
		ResponseErrorWithMsg(c, CodeServerBusy, "获取失败")
		return
	}
	ResponseSuccess(c, gin.H{
		"username":     p1.Username,
		"token":        p1.Token,
		"password":     p1.Password,
		"phone_number": p1.PhoneNumber,
		"email":        p1.Email,
		"is_delete":    p1.IsDelete,
		"identity":     p1.Identity,
		"is_login":     p1.IsLogin,
	})
}

// UpdateMsg 更新信息
func UpdateMsg(c *gin.Context) {
	p := new(models.ParamsUpdateMsg)
	err := c.ShouldBind(&p)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		zap.L().Error("UpdateMsg ShouldBind Error")
		return
	}
	err = service.UpdateMsg(p)
	if err != nil {
		zap.L().Error("service.UpdateMsg error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "修改成功")
}
func GetPictureByFileName(c *gin.Context){
	img:=c.Param("img")
	if img==""{
		ResponseError(c,CodeInvalidParams)
		zap.L().Error("GetFile Vaild param")
		return
	}
	c.File(RootPath+img)
}