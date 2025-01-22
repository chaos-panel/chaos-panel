// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-22 09:27:19
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileBuckets is the golang structure of table chaosplus_file_buckets for DAO operations like Where/Data.
type FileBuckets struct {
	g.Meta    `orm:"table:chaosplus_file_buckets, do:true"`
	Id        interface{} // ID
	TenantId  interface{} // tenant id
	Name      interface{} // name
	Capacity  interface{} // capacity
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // updated at
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // deleted at
}
