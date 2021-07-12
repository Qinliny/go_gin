package library

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

// 数据库信息
type Database struct {
	DbType 		string `json:"db_type"`		// 数据库类型
	DbName 		string `json:"db_name"` 	// 数据库名称
	DbHosts 	string `json:"db_hosts"`	// 数据库地址
	DbPort 		string `json:"db_port"`		// 数据库端口号
	DbUser 		string `json:"db_user"`		// 数据库用户名
	DbPassword	string `json:"db_password"`	// 数据库密码
	DbCharset 	string `json:"db_charset"`	// 数据库字符集
}

func (dbObj Database) Db() *gorm.DB{
	config, err := config.ParseConfig()

	if err != nil {
		panic(err.Error())
	}

	// 数据库连接信息
	dsn := strings.Join([]string{
		config.Database.DbUser, ":", config.Database.DbPassword,
		"@tcp(", config.Database.DbHosts, ":", config.Database.DbPort, ")/",
		config.Database.DbName, "?charset=", config.Database.DbCharset,
	}, "")

	db, err := gorm.Open("mysql" ,dsn)

	if err != nil {
		panic(err)
	}

	// 设置连接池，空闲连接
	db.DB().SetMaxIdleConns(50)
	// 打开链接
	db.DB().SetMaxOpenConns(100)
	// 表明禁用后缀加s
	db.SingularTable(true)

	return db
}