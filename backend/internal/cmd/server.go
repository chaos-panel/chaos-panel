package cmd

import (
	"context"
	"runtime"

	"github.com/chaos-plus/chaos-plus/internal/controller/apis"
	"github.com/chaos-plus/chaos-plus/internal/controller/logs"
	"github.com/chaos-plus/chaos-plus/internal/controller/terminals"
	"github.com/chaos-plus/chaos-plus/internal/controller/users"
	"github.com/chaos-plus/chaos-plus/internal/service"
	"github.com/chaos-plus/chaos-plus/utility/configs"
	"github.com/chaos-plus/chaos-plus/utility/docs"
	"github.com/chaos-plus/chaos-plus/utility/middleware"
	"github.com/chaos-plus/chaos-plus/utility/migration"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var ServerCmd = &gcmd.Command{
	Name:  "server",
	Usage: "server",
	Brief: "start server",
	Func: func(ctx context.Context, parser *gcmd.Parser) error {

		runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
		runtime.SetBlockProfileRate(1)     // (非必需)开启对阻塞操作的跟踪

		// set global timezone
		timezone := configs.GetConfig(ctx, "server.timezone")
		if timezone == nil || timezone.IsNil() || gconv.String(timezone) == "" {
			timezone = gvar.New("UTC")
		}
		println("set timezone: " + gconv.String(timezone))
		if err := gtime.SetTimeZone(gconv.String(timezone)); err != nil {
			panic(err)
		}

		gres.Dump()

		migration.Migrate(ctx)

		s := g.Server()

		s.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(ghttp.MiddlewareHandlerResponse)
			group.Middleware(middleware.ApiRequestNoCacheHandler)
			group.Middleware(middleware.ApiRequestDefaultCorsHandler)
			group.Middleware(middleware.ApiRequestLogsHandler)
			group.Bind(
				apis.NewV1(),
				logs.NewV1(),
				users.NewV1(),
				terminals.NewV1(),
			)
		})

		docs.InitSwagger(ctx, s)

		service.WorkerId().InitOrPanic(ctx, s.GetListenedAddress(), s.GetListenedPorts())
		service.Guid().InitOrPanic(ctx)

		// s.EnablePProf()
		// s.SetServerRoot("resource/public")

		s.Run()
		return nil
	},
}

func init() {
	Main.AddCommand(ServerCmd)
}
