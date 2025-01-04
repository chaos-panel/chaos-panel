package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gcmd"
)

var VersionCmd = &gcmd.Command{
	Name:  "version",
	Usage: "version",
	Brief: "show version info",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		println(gbuild.Info().Version + "@" + gbuild.Info().Time)
		return nil
	},
}

func init() {
	Main.AddCommand(VersionCmd)
}
