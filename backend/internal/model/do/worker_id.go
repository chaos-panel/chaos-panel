// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WorkerId is the golang structure of table chaosplus_worker_id for DAO operations like Where/Data.
type WorkerId struct {
	g.Meta      `orm:"table:chaosplus_worker_id, do:true"`
	Id          interface{} // id
	HostInfo    interface{} // host info
	HostTag     interface{} // host tag
	TimeCreated *gtime.Time // time created
	TimeUpdated *gtime.Time // time updated
}
