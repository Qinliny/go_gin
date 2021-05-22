package model

import (
	"GinDemo/library"
)

var db = library.Db()

type IndexModel struct {
	Id			int64	`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	CreateTime	string	`json:"create_time"`
}

// 初始化表名
func (IndexModel) TableName() string {
	return "users"
}

// 获取用户列表
func (index *IndexModel) GetUserList() []IndexModel{
	var result []IndexModel
	db.Find(&result)
	return result
}

// 根据用户ID获取用户数据
func (index IndexModel) GetUserInfoById(userId string) IndexModel{
	db.Where("id = ?", userId).First(&index)
	return index
}