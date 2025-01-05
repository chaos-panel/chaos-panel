package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var Usage = `
Usage:
	chaosplus [command] [arguments]

The commands are:
	server 		启动服务
	update		更新
	backup		备份
	restore		恢复
	version		版本信息

Examples:
	chaosplus server
	chaosplus update
	chaosplus backup
	chaosplus restore
	chaosplus version
`

var Main = gcmd.Command{
	Name:  "chaosplus",
	Usage: "chaosplus",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		println(Usage)
		return nil
	},
}
