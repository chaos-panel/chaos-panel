package logs

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/logs/v1"
	"github.com/chaos-plus/chaos-plus/internal/service"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return service.Logs().GetList(ctx, req)
}
