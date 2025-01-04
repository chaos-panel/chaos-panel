package cmd

import (
	"context"
	"runtime"

	"github.com/chaos-panel/chaos-panel/internal/controller/logs"
	"github.com/chaos-panel/chaos-panel/utility/docs"
	"github.com/chaos-panel/chaos-panel/utility/middleware"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var ServerCmd = &gcmd.Command{
	Name:  "server",
	Usage: "server",
	Brief: "start server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

		runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // (非必需)开启对阻塞操作的跟踪

		s := g.Server()
		s.EnablePProf()

		docs.InitSwagger(ctx, s)

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareHandlerResponse)
			group.Middleware(middleware.ApiRequestNoCacheHandler)
			group.Middleware(middleware.ApiRequestDefaultCorsHandler)
			group.Middleware(middleware.ApiRequestLogsHandler)
			group.Bind(
				logs.NewV1(),
			)
		})
		s.Group("/tpl", func(group *ghttp.RouterGroup) {
			group.GET("/template", func(r *ghttp.Request) {
				r.Response.WriteTplDefault(g.Map{
					"name": "GoFrame",
				})
			})
		})
		s.Run()
		return nil
	},
}

func init() {
	Main.AddCommand(ServerCmd)
}
