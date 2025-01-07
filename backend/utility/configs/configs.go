package configs

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
)

func GetConfigOrDefault(ctx context.Context, key string, defaultValue interface{}) *gvar.Var {
	if val := GetConfig(ctx, key); val != nil {
		return val
	}
	return gvar.New(defaultValue)
}

func GetConfig(ctx context.Context, key string) *gvar.Var {
	val, err := g.Cfg().GetWithCmd(ctx, key)
	if err != nil {
		panic(err)
	}
	if val != nil {
		return val
	}
	val, err = g.Cfg().GetWithEnv(ctx, key)
	if err != nil {
		panic(err)
	}
	return val
}
