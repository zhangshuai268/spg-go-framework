package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

var Conf *Config

func InitConfig() (*Config, error) {

	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	config := new(Config)

	path := filepath.Join(dir, "internal/config/config.json")
	if os.Getenv("STAGE") != "" {
		path = filepath.Join(dir, "internal/config/config_"+os.Getenv("STAGE")+".json")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("打开配置文件错误" + path + err.Error())
	}

	confByte, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("读取配置文件错误" + err.Error())
	}

	err = json.Unmarshal(confByte, config)
	if err != nil {
		return nil, errors.New("读取配置文件错误" + err.Error())
	}

	Conf = config

	return Conf, nil
}
