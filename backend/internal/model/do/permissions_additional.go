// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PermissionsAdditional is the golang structure of table chaosplus_permissions_additional for DAO operations like Where/Data.
type PermissionsAdditional struct {
	g.Meta    `orm:"table:chaosplus_permissions_additional, do:true"`
	Id        interface{} // ID
	RoleId    interface{} // role id
	Type      interface{} // type
	Value     interface{} // value
	Effect    interface{} // effect
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // updated at
}
