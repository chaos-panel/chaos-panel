// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Plugins is the golang structure of table chaosplus_plugins for DAO operations like Where/Data.
type Plugins struct {
	g.Meta      `orm:"table:chaosplus_plugins, do:true"`
	Id          interface{} // ID
	TenantId    interface{} // tenant id
	RepoId      interface{} // repo id
	Kind        interface{} // kind
	Type        interface{} // type
	Name        interface{} // name
	Author      interface{} // author
	VersionCode interface{} // version code
	VersionName interface{} // version name
	Desc        interface{} // desc
	CreatedBy   interface{} // created by
	CreatedAt   *gtime.Time // created at
	UpdatedBy   interface{} // updated by
	UpdatedAt   *gtime.Time // updated at
	DeletedBy   interface{} // deleted by
	DeletedAt   *gtime.Time // deleted at
}
