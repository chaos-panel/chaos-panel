package configs

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

func GetConfig(ctx context.Context, key string) interface{} {
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
