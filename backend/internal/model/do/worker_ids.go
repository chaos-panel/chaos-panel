// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-31 18:32:57
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerIds is the golang structure of table chaosplus_worker_ids for DAO operations like Where/Data.
type WorkerIds struct {
	g.Meta    `orm:"table:chaosplus_worker_ids, do:true"`
	Id        interface{} // id
	HostInfo  interface{} // host info
	ExpiredAt *gtime.Time // expired at
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // time updated
}
