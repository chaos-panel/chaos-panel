package v1

import (
	"time"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/frame/g"
)

type GetOneReq struct {
	g.Meta `path:"/plugins/:id" tags:"Plugins" method:"get" summary:"get a plugin info"`
	Id     int64 `p:"id" in:"path"`
}
type GetOneRes entity.Plugins

type GetListFilter struct {
	CreatedAt time.Time `p:"createdAt" field:"created_at" table:""`
	Host      string    `p:"host" field:"host" table:""`
}
type GetListOrder struct {
	CreatedAt string `p:"createdAt" field:"created_at" table:""`
}
type GetListReq struct {
	page.PageReq[GetListFilter, GetListOrder]
	g.Meta `path:"/plugins" tags:"Plugins" method:"get" summary:"get a plugin list"`
}
type GetListRes struct {
	page.PageRes[[]entity.Plugins]
}
