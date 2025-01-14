// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package apis

import (
	"context"

	"github.com/chaos-plus/chaos-plus/api/apis/v1"
)

type IApisV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
}
