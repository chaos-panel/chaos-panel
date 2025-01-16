package v1

import (
	"time"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"/terminals" tags:"Terminals" method:"post" summary:"create a new terminal"`
	Type     string `p:"type"       v:"required"`
	Host     string `p:"host"       v:"required"`
	Port     string `p:"port"       v:"required"`
	Username string `p:"username"   v:"required"`
	Password string `p:"password"   v:"required"`
}
type CreateRes entity.Terminals

type GetOneReq struct {
	g.Meta `path:"/terminals/:id" tags:"Terminals" method:"get" summary:"get a terminal info"`
	Id     int64 `p:"id" in:"path"`
}
type GetOneRes entity.Terminals

type GetListFilter struct {
	CreatedAt time.Time `p:"createdAt" field:"created_at" table:""`
	Host      string    `p:"host" field:"host" table:""`
}
type GetListOrder struct {
	CreatedAt string `p:"createdAt" field:"created_at" table:""`
}
type GetListReq struct {
	page.PageReq[GetListFilter, GetListOrder]
	g.Meta `path:"/terminals" tags:"Terminals" method:"get" summary:"get a terminal list"`
}
type GetListRes struct {
	page.PageRes[[]entity.Terminals]
}

type DeleteReq struct {
	g.Meta `path:"/terminals/:id" tags:"Terminals" method:"delete" summary:"delete a terminal"`
	Id     int64 `p:"id"  v:"required#missing id"  in:"path"`
}
type DeleteRes int64
