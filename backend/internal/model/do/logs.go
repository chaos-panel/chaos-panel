// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-06 00:06:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure of table chaosplus_logs for DAO operations like Where/Data.
type Logs struct {
	g.Meta    `orm:"table:chaosplus_logs, do:true"`
	Id        interface{} // id
	Status    interface{} // status
	Log       interface{} // log
	CreatedBy interface{} // created by
	CreatedAt *gtime.Time // created at
}
