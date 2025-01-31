// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UsersRoles is the golang structure of table chaosplus_users_roles for DAO operations like Where/Data.
type UsersRoles struct {
	g.Meta    `orm:"table:chaosplus_users_roles, do:true"`
	Id        interface{} // ID
	UserId    interface{} // user id
	RoleId    interface{} // role id
	ValidedAt *gtime.Time // valided at
	ExpiredAt *gtime.Time // expired at
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
}
