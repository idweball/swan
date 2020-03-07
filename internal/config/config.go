package config

import (
	"github.com/BurntSushi/toml"
	"swan/internal/template"
)

//Storage 后端存储的相关配置
type Storage struct {
	//Type 存储类型
	Type string `toml:"type"`
	//Config 在根据Type创建Storage时，会将Config传递Storage的构造函数
	Config map[string]interface{} `toml:"config"`
}

type Log struct {
	Level string `toml:"level"`
}

//Config 配置文件内容
type Config struct {
	Storage   Storage           `toml:"storage"`
	Templates []template.Config `toml:"templates"`
	Log       Log               `toml:"log"`
}

//CfgParser 解析配置文件
func CfgParser(filename string) (cfg *Config, err error) {
	_, err = toml.DecodeFile(filename, &cfg)
	return
}