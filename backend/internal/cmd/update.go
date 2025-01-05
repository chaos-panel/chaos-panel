package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var UpdateCmd = &gcmd.Command{
	Name:  "update",
	Usage: "update",
	Brief: "update chaosplus version",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}

func init() {
	Main.AddCommand(UpdateCmd)
}
