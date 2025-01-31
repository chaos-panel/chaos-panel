// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table chaosplus_users for DAO operations like Where/Data.
type Users struct {
	g.Meta       `orm:"table:chaosplus_users, do:true"`
	Id           interface{} // ID
	TenantId     interface{} // tenant id
	Username     interface{} // username
	Firstname    interface{} // firstname
	Lastname     interface{} // lastname
	Nickname     interface{} // nickname
	Password     interface{} // password
	Phone        interface{} // phone
	Email        interface{} // email
	AvatarFileId interface{} // avatar file id
	WxUnionId    interface{} // wx union id
	WxmaOpenId   interface{} // wxma open id
	WxmpOpenId   interface{} // wxmp open id
	GithubUserId interface{} // github user id
	LockedBy     interface{} // locked by
	LockedAt     *gtime.Time // locked at
	CreatedBy    interface{} // created by
	CreatedAt    *gtime.Time // created at
	UpdatedBy    interface{} // updated by
	UpdatedAt    *gtime.Time // updated at
	DeletedBy    interface{} // deleted by
	DeletedAt    *gtime.Time // deleted at
}
