package plugin

import (
	"github.com/traefik/yaegi/interp"
)

// GoPlugin yaegi插件实现
type GoPlugin struct {
	interpreter *interp.Interpreter
}

// NewGoPlugin 创建Go插件
func NewGoPlugin() *GoPlugin {
	return &GoPlugin{
		interpreter: interp.New(interp.Options{}),
	}
}
