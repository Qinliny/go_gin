package server

import (
	"Qinly/app/dao"
	"Qinly/library/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// 登录的请求参数
type CheckLoginReq struct {
	Username   string `json:"username" validate:"required,min=8,max=18,email"`
	Password   string `json:"password" validate:"required"`
	VerifyCode string `json:"verify_code" validate:"required"`
}

// 验证登录信息
func CheckLogin(ctx *gin.Context) {
	var param CheckLoginReq
	// 获取请求的参数
	ctx.ShouldBind(&param)
	// 验证请求参数
	if err := checkRequestData(param); err != nil {
		log.ErrorF("验证登录的参数有误，请求参数为: %s, 错误信息：%s", param, err.Error())
		Failed(ctx, "参数有误，请检查后再次提交！")
		return
	}

	verifyCode := "123456"
	// 判断验证码是否正确
	if param.VerifyCode != verifyCode {
		Failed(ctx, "验证码不正确")
		return
	}

	var usersApi dao.UsersApi
	// 判断用户是否存在
	userInfo, err := usersApi.GetUserInfoByUserName(param.Username)
	if err != nil {
		Failed(ctx, err.Error())
		return
	}

	// 判断密码是否正确
	if userInfo.Password != param.Password {
		Failed(ctx, "用户密码不正确，请重新输入！")
		return
	}

	// 加密形成token
	token := EncodeToken(map[string]string{"userId": strconv.FormatInt(userInfo.Id, 10)})
	Success(ctx, []string{token}, "验证通过，允许登录！")
}

// 验证登录请求数据的完整性
func checkRequestData(param CheckLoginReq) error {
	validate := validator.New()
	err := validate.Struct(param)
	if err != nil {
		return err
	}
	return nil
}
