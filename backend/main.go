package main

import (
	_ "chaos-panel/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"chaos-panel/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
