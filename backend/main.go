package main

import (
	"github.com/chaos-plus/chaos-plus/internal/cmd"
	_ "github.com/chaos-plus/chaos-plus/internal/packed"
	_ "github.com/chaos-plus/chaos-plus/utility/migration"

	_ "github.com/chaos-plus/chaos-plus/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
