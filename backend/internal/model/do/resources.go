// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Resources is the golang structure of table chaosplus_resources for DAO operations like Where/Data.
type Resources struct {
	g.Meta        `orm:"table:chaosplus_resources, do:true"`
	Id            interface{} // ID
	TenantId      interface{} // tenant id
	Path          interface{} // path
	PathDepth     interface{} // path depth
	PathMd5       interface{} // path md5
	PathSha1      interface{} // path sha1
	PathSha256    interface{} // path sha256
	PathSha512    interface{} // path sha512
	ChildrenCount interface{} // children count
	CreatedBy     interface{} // created by
	CreatedAt     *gtime.Time // created at
	UpdatedBy     interface{} // updated by
	UpdatedAt     *gtime.Time // updated at
	DeletedBy     interface{} // deleted by
	DeletedAt     *gtime.Time // deleted at
}
