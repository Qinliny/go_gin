package library

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"strings"
)

func Db() *gorm.DB {
	// 数据库连接信息
	dsn := strings.Join([]string{
		ConfigParam.Database.DbUser, ":", ConfigParam.Database.DbPassword,
		"@tcp(", ConfigParam.Database.DbHosts, ":", ConfigParam.Database.DbPort, ")/",
		ConfigParam.Database.DbName, "?charset=", ConfigParam.Database.DbCharset,
	}, "")

	db, err := gorm.Open("mysql", dsn)

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
