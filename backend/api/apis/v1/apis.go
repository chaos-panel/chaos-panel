package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type GetListReq struct {
	g.Meta `path:"/apis" tags:"APIs" method:"get" summary:"get all api list"`
}
type GetListRes struct {
	ApiInfo []ghttp.RouterItem `json:"apiInfo"`
}
