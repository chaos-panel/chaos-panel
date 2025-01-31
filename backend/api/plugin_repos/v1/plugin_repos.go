package v1

import (
	"time"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta   `path:"/plugin-repos" tags:"PluginRepos" method:"post" summary:"create a new repo"`
	Name     string `p:"name"       v:"required"`
	Url      string `p:"url"        v:"required"`
	Desc     string `p:"desc"       v:"required"`
}
type CreateRes entity.PluginRepos

type GetOneReq struct {
	g.Meta `path:"/plugin-repos/:id" tags:"PluginRepos" method:"get" summary:"get a repo info"`
	Id     int64 `p:"id" in:"path"`
}
type GetOneRes entity.PluginRepos

type GetListFilter struct {
	CreatedAt time.Time `p:"createdAt" field:"created_at" table:""`
	Name      string    `p:"name" field:"name" table:""`
}
type GetListOrder struct {
	CreatedAt string `p:"createdAt" field:"created_at" table:""`
}
type GetListReq struct {
	page.PageReq[GetListFilter, GetListOrder]
	g.Meta `path:"/plugin-repos" tags:"PluginRepos" method:"get" summary:"get a repo list"`
}
type GetListRes struct {
	page.PageRes[[]entity.PluginRepos]
}

type DeleteReq struct {
	g.Meta `path:"/plugin-repos/:id" tags:"PluginRepos" method:"delete" summary:"delete a repo"`
	Id     int64 `p:"id"  v:"required#missing id"  in:"path"`
}
type DeleteRes int64
