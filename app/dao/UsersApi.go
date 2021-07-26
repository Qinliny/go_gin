package dao

import (
	"Qinly/app/model"
	"errors"
)

var (
	UsernameNotFound = "用户找不到"
)

type UsersApi struct{}

// 根据用户名获取用户信息
func (users UsersApi) GetUserInfoByUserName(username string) (model.Users, error) {
	usersInfo := new(model.Users).GetUserInfoByUserName(username)
	if usersInfo.Id == 0 {
		return usersInfo, errors.New(UsernameNotFound)
	}
	// 返回数据
	return usersInfo, nil
}
