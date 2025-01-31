package plugin_repos

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/chaos-plus/chaos-plus/api/plugin_repos/v1"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
