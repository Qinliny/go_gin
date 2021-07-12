package library

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	App      App      `json:"app"`
	Database Database `json:"database"`
	Redis	 Redis	  `json:"redis"`
}

// 主要配置信息
type App struct {
	Prot      string `json:"prot"`			// 启动的端口号
	PrivaeKey string `json:"privae_key"`	// 形成token的密钥参数
}


var config *Config = nil

// 解析配置文件
func (configObj Config) ParseConfig() (*Config, error) {
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
