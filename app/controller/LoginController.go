package controller

import (
	"Qinly/app"
	"Qinly/app/api"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 登录的请求参数
type CheckLoginReq struct {
	Username	string	`json:"username" validate:"required,min=8,max=18,email"	`
	Password	string	`json:"password" validate:"required"`
	VieryCode	string	`json:"viery_code validate:"required"`
}

// 登录控制器
type LoginController struct {
	BaseController
	// 登录的请求参数
	CheckLoginReq
}

// 验证登录信息
func (login LoginController)CheckLogin(ctx *gin.Context) {
	// 获取请求的参数
	ctx.ShouldBind(&login.CheckLoginReq)

	// 验证请求参数
	if err := checkRequestData(login.CheckLoginReq); err != nil {
		app.Failed(ctx, err.Error())
		return
	}

	// 赋值请求参数
	param := login.CheckLoginReq

	var usersApi api.UsersApi
	// 判断用户是否存在
	userInfo, err := usersApi.GetUserInfoByUserName(param.Username)
	if err != nil {
		app.Failed(ctx, err.Error())
		return
	}

	// 判断密码是否正确
	if userInfo.Password != param.Password {
		app.Failed(ctx, "用户密码不正确，请重新输入！")
		return
	}
	app.Success(ctx,nil, "验证通过，允许登录！")
}

// 验证登录请求数据的完整性
func checkRequestData(param CheckLoginReq) error{
	validate := validator.New()

	err := validate.Struct(param)
	if err != nil {
		return err
	}
	return nil
}