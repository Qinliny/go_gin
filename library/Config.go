package library

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
	Redis    Redis    `json:"redis"`
}

// 主要配置信息
type App struct {
	Port       string `json:"port"`        // 启动的端口号
	PrivateKey string `json:"private_key"` // 形成token的密钥参数
}

// 数据库信息
type Database struct {
	DbType     string `json:"db_type"`     // 数据库类型
	DbName     string `json:"db_name"`     // 数据库名称
	DbHosts    string `json:"db_hosts"`    // 数据库地址
	DbPort     string `json:"db_port"`     // 数据库端口号
	DbUser     string `json:"db_user"`     // 数据库用户名
	DbPassword string `json:"db_password"` // 数据库密码
	DbCharset  string `json:"db_charset"`  // 数据库字符集
}

// redis配置信息
type Redis struct {
	RedisHost     string `json:"redis_host"`     // redis地址
	RedisPort     string `json:"redis_port"`     // redis端口
	RedisPassword string `json:"redis_password"` // redis密码
	RedisDb       int    `json:"redis_db"`       // redis选择数据库
}

// 配置文件地址
const ConfigJsonPath = "./config/config.json"

var ConfigParam Config

func init() {
	file, err := ioutil.ReadFile(ConfigJsonPath)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &ConfigParam)
	if err != nil {
		fmt.Println(err)
	}
}
