package configs

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
)

func GetConfigOrDefault(ctx context.Context, key string, defaultValue interface{}) *gvar.Var {
	if val := GetConfig(ctx, key); val != nil && !val.IsNil() && !val.IsEmpty() {
		return val
	}
	return gvar.New(defaultValue)
}

func GetConfig(ctx context.Context, key string) *gvar.Var {
	val, err := g.Cfg().GetWithCmd(ctx, key)
	if err != nil {
		panic(err)
	}
	if val != nil && !val.IsNil() && !val.IsEmpty() {
		return val
	}
	val, err = g.Cfg().GetWithEnv(ctx, key)
	if err != nil {
		panic(err)
	}

	if val != nil && !val.IsNil() && !val.IsEmpty() {
		return val
	}

	data := gbuild.Info().Data[key]
	if data != nil {
		return gvar.New(data)
	}
	return gvar.New(nil)
}
