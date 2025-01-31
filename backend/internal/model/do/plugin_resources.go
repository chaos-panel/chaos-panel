// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PluginResources is the golang structure of table chaosplus_plugin_resources for DAO operations like Where/Data.
type PluginResources struct {
	g.Meta      `orm:"table:chaosplus_plugin_resources, do:true"`
	Id          interface{} // ID
	PluginId    interface{} // plugin id
	FileType    interface{} // file type
	FileName    interface{} // file name
	FileId      interface{} // file id
	MountType   interface{} // mount type
	MountPath   interface{} // mount path
	MountPolicy interface{} // mount policy
	CreatedBy   interface{} // created by
	CreatedAt   *gtime.Time // created at
	UpdatedBy   interface{} // updated by
	UpdatedAt   *gtime.Time // updated at
	DeletedBy   interface{} // deleted by
	DeletedAt   *gtime.Time // deleted at
}
