package main

import (
	_ "github.com/chaos-panel/chaos-panel/internal/packed"
	_ "github.com/chaos-panel/chaos-panel/utility/migration"

	"github.com/chaos-panel/chaos-panel/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
