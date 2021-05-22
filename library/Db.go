package library

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

func Db() *gorm.DB{
	config, err := ParseConfig()

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