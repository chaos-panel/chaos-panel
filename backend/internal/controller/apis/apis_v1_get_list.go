package apis

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/apis/v1"
	"github.com/gogf/gf/v2/net/ghttp"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	s := ghttp.GetServer(ghttp.DefaultServerName)
	return &v1.GetListRes{
		ApiInfo: s.GetRoutes(),
	}, nil
}
