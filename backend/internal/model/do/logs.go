// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-16 23:38:52
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure of table chaosplus_logs for DAO operations like Where/Data.
type Logs struct {
	g.Meta     `orm:"table:chaosplus_logs, do:true"`
	Id         interface{} // id
	TenantId   interface{} // tenant id
	ClientType interface{} // client type
	ClientInfo interface{} // client info
	Remark     interface{} // remark
	Log        interface{} // log
	CreatedBy  interface{} // created by
	CreatedAt  *gtime.Time // locked at
}
