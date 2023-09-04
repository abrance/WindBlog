package config

import "gopkg.in/ini.v1"

var DefaultCfg = &Config{}
var defaultConfigPath = "./config.ini"

// TODO 设置各种选项默认值

func LoadDefaultConfig(path string) error {
	var (
		pf  *ini.File
		err error
	)
	pf, err = ini.Load(path)
	if err != nil {
		return err
	}
	err = pf.MapTo(DefaultCfg)
	if err != nil {
		return err
	}
	return nil
}

type Default struct {
	AddrLs []string `ini:"addr_ls"`
}

type Config struct {
	Default `ini:"default"`
}

// Load 只会在启动时被调用
func Load() error {
	err := LoadDefaultConfig(defaultConfigPath)
	if err != nil {
		return err
	}
	return nil
}

// GetConfig 获取配置使用这个方法
func GetConfig() Config {
	return *DefaultCfg
}
