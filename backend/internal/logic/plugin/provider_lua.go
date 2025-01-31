package plugin

import (
	lua "github.com/yuin/gopher-lua"
)

// LuaPlugin lua插件实现
type LuaPlugin struct {
	state *lua.LState
}

// NewLuaPlugin 创建Lua插件
func NewLuaPlugin() *LuaPlugin {
	return &LuaPlugin{
		state: lua.NewState(),
	}
}
