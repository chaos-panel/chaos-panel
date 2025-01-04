package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var RestoreCmd = &gcmd.Command{
	Name:  "restore",
	Usage: "restore",
	Brief: "restore all data from a zip file",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}

func init() {
	Main.AddCommand(RestoreCmd)
}
