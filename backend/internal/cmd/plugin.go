package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var PluginCmd = &gcmd.Command{
	Name:  "plugin",
	Usage: "plugin",
	Brief: "manage plugins",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}

func init() {
	Main.AddCommand(PluginCmd)
}
