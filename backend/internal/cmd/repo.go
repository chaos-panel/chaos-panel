package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var RepoCmd = &gcmd.Command{
	Name:  "repo",
	Usage: "repo",
	Brief: "manage plugin repository",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		return nil
	},
}

func init() {
	Main.AddCommand(RepoCmd)
}
