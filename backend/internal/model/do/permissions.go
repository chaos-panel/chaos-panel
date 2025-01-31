// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure of table chaosplus_permissions for DAO operations like Where/Data.
type Permissions struct {
	g.Meta     `orm:"table:chaosplus_permissions, do:true"`
	Id         interface{} // ID
	RoleId     interface{} // role id
	ResourceId interface{} // resource id
	Action     interface{} // action
	Effect     interface{} // effect
	CreatedBy  interface{} // created by
	CreatedAt  *gtime.Time // created at
	UpdatedBy  interface{} // updated by
	UpdatedAt  *gtime.Time // updated at
}
