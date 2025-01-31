// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PluginRepos is the golang structure of table chaosplus_plugin_repos for DAO operations like Where/Data.
type PluginRepos struct {
	g.Meta    `orm:"table:chaosplus_plugin_repos, do:true"`
	Id        interface{} // ID
	TenantId  interface{} // tenant id
	Name      interface{} // name
	Url       interface{} // url
	Desc      interface{} // desc
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // updated at
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // deleted at
}
