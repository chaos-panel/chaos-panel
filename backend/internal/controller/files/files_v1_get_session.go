package files

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/chaos-plus/chaos-plus/api/files/v1"
)

func (c *ControllerV1) GetSession(ctx context.Context, req *v1.GetSessionReq) (res *v1.GetSessionRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
