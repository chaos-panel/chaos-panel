package plugin

import (
	"github.com/dop251/goja"
)

// JsPlugin javascript插件实现
type JsPlugin struct {
	runtime *goja.Runtime
}

func NewJsPlugin() *JsPlugin {
	return &JsPlugin{
		runtime: goja.New(),
	}
}
