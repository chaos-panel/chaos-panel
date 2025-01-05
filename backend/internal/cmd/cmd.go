package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var Main = gcmd.Command{
	Name:  "chaosplus",
	Usage: "chaosplus",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}
