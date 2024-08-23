package global

import (
	"github.com/kulisi/global/config"
	"github.com/kulisi/global/internal"
	"github.com/kulisi/global/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"path/filepath"
)

func InitializeZapLogger(conf config.Zap) *zap.Logger {
	if ok, _ := utils.PathExists(filepath.Join(utils.ExecPath(), "log")); !ok {
		_ = os.Mkdir(path.Join(utils.ExecPath(), "log"), os.ModePerm)
	}
	levels := conf.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(&internal.ZapCoreConfig{
			Level:        levels[i],
			Encoder:      conf.Encoder(),
			Director:     conf.Director,
			RetentionDay: conf.RetentionDay,
			LogInConsole: conf.LogInConsole,
		})
		cores = append(cores, core)
	}
	logger := zap.New(zapcore.NewTee(cores...))
	if conf.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
