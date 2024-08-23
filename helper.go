package global

import (
	"encoding/json"
	"github.com/kulisi/global/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"sync"
)

var (
	g    *Global
	once sync.Once
)

func Setup(config *config.Config) {
	once.Do(func() {
		g = NewGlobalByConfig(config)
	})
}

// Debug 输出 Debug 级别日志
func Debug(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Debug(msg, fields...)
	}
}

// Info 输出 Info 级别日志
func Info(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Info(msg, fields...)
	}
}

// Warn 输出 Warn 级别日志
func Warn(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Warn(msg, fields...)
	}
}

// Error 输出 Error 级别日志
func Error(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Error(msg, fields...)
	}
}

// DPanic 输出 DPanic 级别日志
func DPanic(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.DPanic(msg, fields...)
	}
}

// Panic 输出 Panic 级别日志
func Panic(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Panic(msg, fields...)
	}
}

// Fatal 输出 Fatal 级别日志
func Fatal(msg string, fields ...zap.Field) {
	if g.conf.Zap.Use {
		g.log.Fatal(msg, fields...)
	}
}

// WatchConfig 监听配置文件
func WatchConfig(c ConfigChange) {
	g.v.WatchConfig()
	g.v.OnConfigChange(c)
}

// DB 返回数据库连接实例
func DB() *gorm.DB {
	return g.db
}

// Config 返回 config 实例
func Config() *config.Config {
	return g.conf
}

// GetConfigString 返回 config json
func GetConfigString() string {
	js, _ := json.MarshalIndent(g.conf, "", "\t")
	return string(js)
}
