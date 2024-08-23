package global

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/kulisi/global/config"
	"github.com/kulisi/global/internal"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
	"strings"
)

// ConfigChange 监听配置文件发生变动时 的回调函数
type ConfigChange func(event fsnotify.Event)

type Global struct {
	v    *viper.Viper
	conf *config.Config
	log  *zap.Logger
	db   *gorm.DB
}

func NewGlobalByConfig(config *config.Config) *Global {
	v := viper.New()
	v.SetConfigName(config.Viper.ConfigName)
	v.SetConfigType(config.Viper.ConfigType)
	for _, path := range config.Viper.ConfigPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		log.Fatalln(fmt.Sprintf("read in config failed.\n%s\n", err.Error()))
	}
	if err := v.Unmarshal(config); err != nil {
		log.Fatalln(fmt.Sprintf("unmarshal config failed.\n%s\n", err.Error()))
	}

	// 创建日志记录实例
	var l *zap.Logger
	if config.Zap.Use {
		l = InitializeZapLogger(config.Zap)
	}

	// 创建数据库实例
	var db *gorm.DB
	switch strings.ToUpper(config.Gorm.Use) {
	case "mssql":
		db = internal.Mssql(config.Gorm.Mssql)
	case "mysql":
		db = internal.Mysql(config.Gorm.Mysql)
	}

	return &Global{
		v:    v,
		conf: config,
		log:  l,
		db:   db,
	}
}
