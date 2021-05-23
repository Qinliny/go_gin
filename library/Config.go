package library

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
}

// 主要配置信息
type App struct {
	Prot      string `json:"prot"`
	PrivaeKey string `json:"privae_key"`
}

// 数据库信息
type Database struct {
	// 数据库类型
	DbType string `json:"db_type"`
	// 数据库名称
	DbName string `json:"db_name"`
	// 数据库地址
	DbHosts string `json:"db_hosts"`
	// 数据库端口号
	DbPort string `json:"db_port"`
	// 数据库用户名
	DbUser string `json:"db_user"`
	// 数据库密码
	DbPassword string `json:"db_password"`
	// 数据库字符集
	DbCharset string `json:"db_charset"`
}

var config *Config = nil

// 解析配置文件
func ParseConfig() (*Config, error) {
	// 根据传入的地址 打开配置文件
	file, err := os.Open(CONFIG_JSON_PATH)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	// 传入的是一个io文件
	reader := bufio.NewReader(file)
	decoer := json.NewDecoder(reader)
	if err := decoer.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
