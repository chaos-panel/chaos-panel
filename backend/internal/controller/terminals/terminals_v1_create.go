package terminals

import (
	"context"

	v1 "github.com/chaos-plus/chaos-plus/api/terminals/v1"
	"github.com/chaos-plus/chaos-plus/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return service.Terminal().Create(ctx, req)
}
