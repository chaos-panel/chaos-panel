// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-18 00:02:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DistributedLocks is the golang structure of table chaosplus_distributed_locks for DAO operations like Where/Data.
type DistributedLocks struct {
	g.Meta    `orm:"table:chaosplus_distributed_locks, do:true"`
	LockKey   interface{} // lock key
	HostInfo  interface{} // host info
	ExpiredAt *gtime.Time // expired at
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	UpdatedBy interface{} // updated by
	UpdatedAt *gtime.Time // time updated
}
