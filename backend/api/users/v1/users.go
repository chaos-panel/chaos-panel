package v1

import (
	"time"

	"github.com/chaos-plus/chaos-plus/internal/model/entity"
	"github.com/chaos-plus/chaos-plus/utility/page"
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta    `path:"/users" tags:"Users" method:"post" summary:"create a new user"`
	Username  string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Nickname  string `p:"nickname"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Password1 string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Password2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"`
	Phone     string `p:"phone"`
}
type CreateRes entity.Users

type UpdateReq struct {
	g.Meta    `path:"/users/:id" tags:"Users" method:"patch" summary:"update a user info"`
	Id        int64  `p:"id"  v:"required#请输入ID"  in:"path"`
	Username  string `p:"username"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Nickname  string `p:"nickname"  v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Password1 string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Password2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|密码长度不够|两次密码不一致"`
	Phone     string `p:"phone"`
}
type UpdateRes entity.Users

type GetOneReq struct {
	g.Meta `path:"/users/:id" tags:"Users" method:"get" summary:"get a user info"`
	Id     int64 `p:"id"  v:"required#请输入ID"  in:"path"`
}
type GetOneRes entity.Users

type GetListFilter struct {
	CreatedAt time.Time `p:"createdAt" field:"created_at" table:""`
}
type GetListOrder struct {
	CreatedAt string `p:"createdAt" field:"created_at" table:""`
}
type GetListReq struct {
	page.PageReq[GetListFilter, GetListOrder]
	g.Meta `path:"/users" tags:"Users" method:"get" summary:"get a user list"`
}
type GetListRes struct {
	page.PageRes[[]entity.Users]
}

type DeleteReq struct {
	g.Meta `path:"/users/:id" tags:"Users" method:"delete" summary:"delete a user"`
	Id     int64 `p:"id"  v:"required#请输入ID"  in:"path"`
}
type DeleteRes int64

// --------------------------------------------

type GetInfoReq struct {
	g.Meta `path:"/userinfo" tags:"Users" method:"get" summary:"get auth user info"`
}
type GetInfoRes entity.Users
