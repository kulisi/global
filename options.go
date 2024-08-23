package global

import (
	"github.com/kulisi/global/config"
	"github.com/kulisi/global/utils"
)

type Option func(*config.Config)

func DefaultConfig(opts ...Option) *config.Config {
	// 定义默认值
	_config := &config.Config{
		Viper: config.Viper{
			ConfigName:  "config",
			ConfigType:  "yaml",
			ConfigPaths: []string{utils.ExecPath()},
		},
	}
	// 处理选项值转成配置值
	for _, opt := range opts {
		opt(_config)
	}
	// 返回配置结构体
	return _config
}

func ConfigName(in string) Option {
	return func(c *config.Config) {
		if in == "" {
			c.Viper.ConfigName = "config"
		} else {
			c.Viper.ConfigName = in
		}
	}
}

func ConfigType(in string) Option {
	return func(c *config.Config) {
		if in == "" {
			c.Viper.ConfigType = "yaml"
		} else {
			c.Viper.ConfigType = in
		}
	}
}

func AddPath(in string) Option {
	return func(c *config.Config) {
		if in != "" {
			c.Viper.ConfigPaths = append(c.Viper.ConfigPaths, in)
		}
	}
}
