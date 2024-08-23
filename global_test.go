package global

import (
	"fmt"
	"testing"
)

func TestGlobal(t *testing.T) {
	conf := DefaultConfig(AddPath("./"))
	Setup(conf)
	fmt.Println(GetConfigString())
}
