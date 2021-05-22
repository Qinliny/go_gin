package api

import "GinDemo/app/model"

type IndexApi struct {

}

// 获取用户列表数据
func (index *IndexApi) GetUserList() []model.IndexModel{
	result := new(model.IndexModel).GetUserList()
	return result
}

// 根据用户ID获取用户信息
func (index *IndexApi) GetUserInfoById(userId string) model.IndexModel{
	userInfo := new(model.IndexModel).GetUserInfoById(userId)
	return userInfo
}