package global

import (
	"testing"
)

func TestGlobal(t *testing.T) {
	conf := DefaultConfig(AddPath("./"), ConfigName("config_debug"))
	Setup(conf)
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	DPanic("dPanic")
	Panic("panic")
	Fatal("fatal")
}
