// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ChangeLogs is the golang structure of table chaosplus_change_logs for DAO operations like Where/Data.
type ChangeLogs struct {
	g.Meta    `orm:"table:chaosplus_change_logs, do:true"`
	Id        interface{} // id
	UnionId   interface{} // union id
	Snapshot  interface{} // snapshot
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // locked at
	DeletedBy interface{} // deleted by
	DeletedAt *gtime.Time // time deleted
}
