package logs

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/logs/v1"
	"github.com/chaos-plus/chaos-plus/internal/service"
)

func (c *ControllerV1) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	return service.Logs().GetOne(ctx, req)
}
