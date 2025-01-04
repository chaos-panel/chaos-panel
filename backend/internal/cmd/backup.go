package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var BackupCmd = &gcmd.Command{
	Name:  "backup",
	Usage: "backup",
	Brief: "backup all data to a zip file",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}

func init() {
	Main.AddCommand(BackupCmd)
}
