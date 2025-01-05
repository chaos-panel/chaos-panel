// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2025-01-05 22:44:22
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Logs is the golang structure of table chaosplus_logs for DAO operations like Where/Data.
type Logs struct {
	g.Meta    `orm:"table:chaosplus_logs, do:true"`
	Id        interface{} // ID
	Status    interface{} // 状态
	Log       interface{} // 日志
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建时间
}
