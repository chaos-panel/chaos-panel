package v1

import (
	"time"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/frame/g"
)

type GetOneReq struct {
	g.Meta `path:"/logs/:id" tags:"Logs" method:"get" summary:"get a log info"`
	Id     int64 `p:"id"  v:"required#请输入ID" in:"path"`
}
type GetOneRes entity.Logs

type GetListFilter struct {
	CreatedAt time.Time `p:"createdAt" field:"created_at" table:""`
}
type GetListOrder struct {
	CreatedAt string `p:"createdAt" field:"created_at" table:""`
}
type GetListReq struct {
	page.PageReq[GetListFilter, GetListOrder]
	g.Meta `path:"/logs" tags:"Logs" method:"get" summary:"get a log list"`
}
type GetListRes struct {
	page.PageRes[[]entity.Logs]
}
