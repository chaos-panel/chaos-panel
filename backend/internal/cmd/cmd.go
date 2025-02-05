package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
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
	repo		仓库
	plugin		插件
	version		版本信息

Examples:
	chaosplus server
	chaosplus update
	chaosplus backup -o <output>
	chaosplus restore -i <input>
	chaosplus repo <add|rm|list>
	chaosplus plugin <add|rm|list>
	chaosplus version
`

var Main = gcmd.Command{
	Name:  "chaosplus",
	Usage: "chaosplus",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		g.Log().Debug(ctx, gcmd.GetArgAll(), gcmd.GetOptAll(), Usage)
		println(Usage)
		return nil
	},
}
