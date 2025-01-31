package plugin

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// WasmPlugin wasm插件实现
type WasmPlugin struct {
	ctx     context.Context
	runtime wazero.Runtime
	module  api.Module
}

// NewWasmPlugin 创建WASM插件
func NewWasmPlugin() *WasmPlugin {
	return &WasmPlugin{}
}

func (p *WasmPlugin) Init(ctx context.Context, source []byte) (err error) {
	runtime := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter())
	module, err := runtime.Instantiate(p.ctx, source)
	if err != nil {
		return err
	}
	p.ctx = ctx
	p.runtime = runtime
	p.module = module
	return
}
