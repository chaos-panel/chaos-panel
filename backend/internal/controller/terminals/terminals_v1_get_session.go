package terminals

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/terminals/v1"
	"github.com/chaos-plus/chaos-plus/internal/service"
)

func (c *ControllerV1) GetSession(ctx context.Context, req *v1.GetSessionReq) (res *v1.GetSessionRes, err error) {
	return service.Terminal().GetSession(ctx, req)
}
