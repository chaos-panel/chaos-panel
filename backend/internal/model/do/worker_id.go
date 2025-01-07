// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-07 10:37:30
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerId is the golang structure of table chaosplus_worker_id for DAO operations like Where/Data.
type WorkerId struct {
	g.Meta    `orm:"table:chaosplus_worker_id, do:true"`
	Id        interface{} // id
	HostInfo  interface{} // host info
	ExpiredAt *gtime.Time // expired at
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // time updated
}
