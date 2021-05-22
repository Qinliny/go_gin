package model

type Users struct {
	Id			int64	`json:"id"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	CreateTime	string	`json:"create_time"`
}

// 初始化表名
func (Users) TableName() string {
	return "users"
}

// 获取用户列表
func (users *Users) GetUserList() []Users{
	var result []Users
	db.Find(&result)
	return result
}

// 根据用户名获取用户信息
func (users *Users) GetUserInfoByUserName(username string) Users{
	var userInfo Users
	db.Where("username = ?", username).First(&userInfo)
	return userInfo
}